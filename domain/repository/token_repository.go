package repository

import "contractor_panel/domain/model"

type TokenRepository interface {
	SetTokenDetails(userid int64, td *model.TokenDetails) error
	DeleteAuth(givenUuid string) (int64, error)
	FetchAuth(authD *model.AccessDetails) (int64, error)
}
