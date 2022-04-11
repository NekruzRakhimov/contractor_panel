package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Transactional interface {
	WithTransaction(ctx context.Context) (pgx.Tx, error)
	RollbackQuietly(tx pgx.Tx, ctx context.Context)
}
