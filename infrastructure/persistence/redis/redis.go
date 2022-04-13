package redis

import (
	"contractor_panel/application/cerrors"
	"contractor_panel/application/config"
	"crypto/tls"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func RedisConn() *redis.Client {
	dsn := viper.GetString(config.RedisDsn)

	client := redis.NewClient(&redis.Options{
		Addr:      dsn, //redis port
		TLSConfig: &tls.Config{},
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(cerrors.ErrCouldNotConnectToDb(err))
	}

	return client
}
