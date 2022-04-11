package postgres

import "github.com/jackc/pgx/v4/pgxpool"

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db}
}

func (b *UserRepository) FindUserRoles(userId int64) ([]string, error) {
	return nil, nil
}
