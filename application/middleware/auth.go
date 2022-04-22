package middleware

import (
	"context"
	"contractor_panel/application/cerrors"
	"contractor_panel/application/config"
	"contractor_panel/application/respond"
	"contractor_panel/domain/model"
	"contractor_panel/domain/repository"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"strings"
)

const (
	UserInfoCtxKey = "UserInfo"
)

type authenticationHandler struct {
	r    repository.UserRepository
	next http.Handler
}

func newAuthenticationHandler(r repository.UserRepository, next http.Handler) *authenticationHandler {
	return &authenticationHandler{r, next}
}

func (a *authenticationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	accessDetails, err := ExtractTokenData(r)
	if err != nil {
		respond.WithError(w, r, cerrors.ErrCouldNotVerifyToken(err))
		return
	}

	var userInfo *model.UserInfo = nil
	if accessDetails != nil {
		roles, err := a.r.FindUserRoles(accessDetails.UserId)
		if err != nil {
			log.Warn(err)
		}

		userRoles := make([]model.RoleCode, len(roles))
		for i, role := range roles {
			userRoles[i] = model.RoleCode(role)
		}
		userInfo = model.NewUserInfo(accessDetails.UserId, userRoles)
	} else {
		respond.WithError(w, r, cerrors.ErrCouldNotVerifyToken(errors.New("нет данных для доступа")))
		return
	}

	ctx := context.WithValue(r.Context(), UserInfoCtxKey, userInfo)

	a.next.ServeHTTP(w, r.WithContext(ctx))
}

func AuthHandler(r repository.UserRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return newAuthenticationHandler(r, next)
	}
}

func GetUserInfo(ctx context.Context) *model.UserInfo {
	info, ok := ctx.Value(UserInfoCtxKey).(*model.UserInfo)
	if !ok {
		return nil
	}

	return info
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, "Bearer ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func ExtractTokenFromCookie(r *http.Request) (string, error) { //TODO переделать на это когда фронт разберется работать с куки
	c, err := r.Cookie("accessToken")
	if err != nil {
		return "", err
	}

	return c.Value, err
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	// Get the JWT string
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString(config.AccessSecret)), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractTokenData(r *http.Request) (*model.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &model.AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}

	return nil, err
}
