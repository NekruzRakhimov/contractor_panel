package model

type RoleCode string

type UserInfo struct {
	userId int64
	roles  []RoleCode
}

func (u UserInfo) UserId() int64 {
	return u.userId
}

func (u UserInfo) Roles() []RoleCode {
	return u.roles
}

func NewUserInfo(userId int64, roles []RoleCode) *UserInfo {
	return &UserInfo{
		userId: userId,
		roles:  roles,
	}
}
