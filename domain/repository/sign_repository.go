package repository

import (
	"context"
	"contractor_panel/domain/model"
)

type SignRepository interface {
	FindUserCredentials(ctx context.Context, login string) ([]model.Credentials, error)
}
