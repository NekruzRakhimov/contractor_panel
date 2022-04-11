package postgres

import (
	"context"
	"contractor_panel/application/cerrors"
	"contractor_panel/application/config"
	"contractor_panel/infrastructure/persistence"
	"fmt"
	"github.com/jackc/pgtype"
	shopspring "github.com/jackc/pgtype/ext/shopspring-numeric"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func DBConn() *pgxpool.Pool {
	pgConfig, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s search_path=%s",
			viper.GetString(config.DatasourcesPostgresHost),
			viper.GetString(config.DatasourcesPostgresPort),
			viper.GetString(config.DatasourcesPostgresUser),
			viper.GetString(config.DatasourcesPostgresPassword),
			viper.GetString(config.DatasourcesPostgresDatabase),
			viper.GetString(config.DatasourcesPostgresSchema)),
	)
	if err != nil {
		log.Fatal(cerrors.ErrCouldNotConnectToDb(err))
	}

	pgConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		conn.ConnInfo().RegisterDataType(pgtype.DataType{
			Value: &shopspring.Numeric{},
			Name:  "numeric",
			OID:   pgtype.NumericOID,
		})

		return nil
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), pgConfig)
	if err != nil {
		log.Fatal(cerrors.ErrCouldNotConnectToDb(err))
	}

	return pool
}

type pgExecutor interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

//region PgRowsWrapper
type pgRowsWrapper struct {
	rows pgx.Rows
}

func newDbRowsWrapper(rows pgx.Rows) *pgRowsWrapper {
	return &pgRowsWrapper{rows: rows}
}

func (p pgRowsWrapper) Close() error {
	p.rows.Close()
	return nil
}

func (p pgRowsWrapper) Next() bool {
	return p.rows.Next()
}

func (p pgRowsWrapper) Scan(dest ...interface{}) error {
	return p.rows.Scan(dest...)
}

//endregion

func Query(con pgExecutor, ctx context.Context, query string, placeholders ...interface{}) persistence.DbScanner {
	rows, err := con.Query(ctx, query, placeholders...)
	if err != nil {
		return persistence.NewDbScannerWithError(err)
	}

	return persistence.NewDbScanner(newDbRowsWrapper(rows))
}

func QueryWithMap(con pgExecutor, ctx context.Context, query string, placeholders map[string]interface{}) persistence.DbScanner {
	inlineQuery, inlinePlaceholders, err := InlineNamedPlaceholders(query, placeholders)
	if err != nil {
		return persistence.NewDbScannerWithError(err)
	}

	rows, err := con.Query(ctx, inlineQuery, inlinePlaceholders...)
	if err != nil {
		return persistence.NewDbScannerWithError(err)
	}

	return persistence.NewDbScanner(newDbRowsWrapper(rows))
}

func InlineNamedPlaceholders(query string, placeholders map[string]interface{}) (string, []interface{}, error) {
	return persistence.InlineNamedPlaceholders(renamePlaceholder, query, placeholders)
}

func renamePlaceholder(i uint8) string {
	return fmt.Sprintf("$%d", i+1)
}
