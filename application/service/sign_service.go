package service

import (
	"context"
	"contractor_panel/application/cerrors"
	"contractor_panel/application/config"
	"contractor_panel/domain/model"
	"contractor_panel/domain/repository"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

type SignService interface {
	SignIn(ctx context.Context, incomingCredential model.Credentials) (*model.TokenDetails, error)
	RefreshToken(ctx context.Context, refreshToken string) (*model.TokenDetails, error)
	SignOut(ctx context.Context, refreshToken string) error
}

type signService struct {
	r          repository.SignRepository
	tokenRedis repository.TokenRepository
}

func NewSignService(r repository.SignRepository, tokenRedis repository.TokenRepository) SignService {
	return &signService{r, tokenRedis}
}

func (s *signService) SignIn(ctx context.Context, incomingCredential model.Credentials) (*model.TokenDetails, error) {
	existingCredentials, err := s.r.FindUserCredentials(ctx, incomingCredential.UserLogin)
	if err != nil {
		return nil, cerrors.ErrCouldNotSignIn(err, incomingCredential.UserLogin)
	}

	if existingCredentials == nil {
		return nil, cerrors.ErrCouldNotSignIn(errors.New(fmt.Sprintf(
			"по указанному логину %s нет контрагента в базе", incomingCredential.UserLogin)),
			incomingCredential.UserLogin)
	}

	if len(existingCredentials) > 1 {
		return nil, cerrors.ErrCouldNotSignIn(errors.New(fmt.Sprintf(
			"по указанному логину %s в базе заведено более одного контрагента", incomingCredential.UserLogin)),
			incomingCredential.UserLogin)
	}

	if err = existingCredentials[0].Compare(incomingCredential.UserPassword); err != nil {
		return nil, cerrors.ErrCouldNotSignIn(err, incomingCredential.UserLogin)
	}

	td, err := s.CreateToken(existingCredentials[0].Id)
	if err != nil {
		return nil, cerrors.ErrCouldNotSignIn(err, incomingCredential.UserLogin)
	}

	return td, nil
}

func (s *signService) CreateToken(userid int64) (*model.TokenDetails, error) {
	td := &model.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	//Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(viper.GetString(config.AccessSecret)))
	if err != nil {
		return nil, err
	}

	//Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(viper.GetString(config.RefreshSecret)))
	if err != nil {
		return nil, err
	}

	if err = s.tokenRedis.SetTokenDetails(userid, td); err != nil {
		return nil, err
	}

	return td, nil
}

func (s *signService) RefreshToken(ctx context.Context, refreshToken string) (*model.TokenDetails, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString(config.RefreshSecret)), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}

		//Delete the previous Refresh Token
		deleted, err := s.tokenRedis.DeleteAuth(refreshUuid)
		if err != nil { //if any goes wrong
			return nil, err
		}
		if deleted == 0 {
			return nil, errors.New(fmt.Sprintf("в redis нет указанного рефреш токена %s по  ИД пользователя %d",
				refreshUuid, userId))
		}

		return s.CreateToken(userId)
	}

	return nil, err
}

func (s *signService) SignOut(ctx context.Context, refreshToken string) error {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString(config.RefreshSecret)), nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			return err
		}
		userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return err
		}

		//Delete the previous Refresh Token
		deleted, err := s.tokenRedis.DeleteAuth(refreshUuid)
		if err != nil { //if any goes wrong
			return err
		}
		if deleted == 0 {
			return errors.New(fmt.Sprintf("в redis нет указанного рефреш токена %s по  ИД пользователя %d",
				refreshUuid, userId))
		}

		return nil
	}

	return err
}
