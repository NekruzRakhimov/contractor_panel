package repository

type UserRepository interface {
	FindUserRoles(userId int64) ([]string, error)
}
