package redis

import (
	"contractor_panel/domain/model"
	"github.com/go-redis/redis/v7"
	"strconv"
	"time"
)

type TokenRepository struct {
	client *redis.Client
}

func NewTokenRepository(client *redis.Client) *TokenRepository {
	return &TokenRepository{client}
}

func (r *TokenRepository) SetTokenDetails(userid int64, td *model.TokenDetails) error {
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	err := r.client.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *TokenRepository) DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := r.client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func (r *TokenRepository) FetchAuth(authD *model.AccessDetails) (int64, error) {
	userid, err := r.client.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseInt(userid, 10, 64)
	return userID, nil
}
