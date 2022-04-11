package postgres

import (
	"context"
	"contractor_panel/domain/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SignRepository struct {
	db *pgxpool.Pool
}

func NewSignRepository(db *pgxpool.Pool) *SignRepository {
	return &SignRepository{db}
}

func (r *SignRepository) FindUserCredentials(ctx context.Context, login string) ([]model.Credentials, error) {
	args := make(model.NamedArguments)
	args["login"] = login
	query := `select 	c.id, c.email, c.password
				from 	v_contractors_credentials c
				where 	c.is_delete = false 
						and c.status = 'ACTIVE'
						and c.password is not null
						and c.email = :login`

	res, err := QueryWithMap(r.db, ctx, query, args).ReadAll(model.Credentials{})
	if err != nil {
		return nil, err
	}

	return r.unwrapCredentialsSlice(res), nil
}

func (r *SignRepository) unwrapCredentialsSlice(res interface{}) []model.Credentials {
	if res == nil {
		return nil
	} else {
		return res.([]model.Credentials)
	}
}
