package config

import (
	"contractor_panel/application/cerrors"
	"github.com/spf13/viper"
	"time"
)

const (
	Port                        = "PORT"
	AppName                     = "APP_NAME"
	AppInstance                 = "APP_INSTANCE"
	LogLevel                    = "LOG_LEVEL"
	LogPrettyPrint              = "LOG_PRETTY_PRINT"
	CorsAllowedOrigins          = "CORS_ALLOWED_ORIGINS"
	CorsAllowedMethods          = "CORS_ALLOWED_METHODS"
	CorsAllowedHeaders          = "CORS_ALLOWED_HEADERS"
	HealthcheckTimeout          = "HEALTHCHECK_TIMEOUT"
	HttpRequestTimeout          = "HTTP_REQUEST_TIMEOUT"
	DatasourcesPostgresHost     = "DATASOURCES_POSTGRES_HOST"
	DatasourcesPostgresPort     = "DATASOURCES_POSTGRES_PORT"
	DatasourcesPostgresUser     = "DATASOURCES_POSTGRES_USER"
	DatasourcesPostgresPassword = "DATASOURCES_POSTGRES_PASSWORD"
	DatasourcesPostgresDatabase = "DATASOURCES_POSTGRES_DATABASE"
	DatasourcesPostgresSchema   = "DATASOURCES_POSTGRES_SCHEMA"
	AccessSecret                = "ACCESS_SECRET"
	RefreshSecret               = "REFRESH_SECRET"
	RedisDsn                    = "REDIS_DSN"
)

var EncRegex = `(?m)ENC\((.*)\)`

var RequiredEnvs = []string{
	// Core
	Port,
	AppName,
	AppInstance,
	HealthcheckTimeout,

	// CORS
	CorsAllowedOrigins,
	CorsAllowedMethods,
	CorsAllowedHeaders,

	//Token secret keys
	AccessSecret,
	RefreshSecret,

	// Postgres
	DatasourcesPostgresHost,
	DatasourcesPostgresPort,
	DatasourcesPostgresUser,
	DatasourcesPostgresPassword,
	DatasourcesPostgresDatabase,
	DatasourcesPostgresSchema,

	//Redis
	RedisDsn,
}

type defaultEnvValueGetter = func() interface{}

var DefaultEnvs = map[string]interface{}{
	LogLevel:           "info",
	LogPrettyPrint:     false,
	HttpRequestTimeout: time.Second * 60,
}

// CheckEnv проверяет заданные ENV переменные
// и подставляет `default` значения по необходимости
func CheckEnv() error {
	missingEnvs := make([]string, 0)

	for _, requiredEnvKey := range RequiredEnvs {
		if !viper.IsSet(requiredEnvKey) {
			missingEnvs = append(missingEnvs, requiredEnvKey)
		}
	}

	for defaultEnvKey, defaultEnvValue := range DefaultEnvs {
		if !viper.IsSet(defaultEnvKey) {
			if getDefaultValue, ok := defaultEnvValue.(defaultEnvValueGetter); ok {
				viper.SetDefault(defaultEnvKey, getDefaultValue())
			} else {
				viper.SetDefault(defaultEnvKey, defaultEnvValue)
			}
		}
	}

	if len(missingEnvs) > 0 {
		return cerrors.ErrConfigurationError(missingEnvs)
	}

	return nil
}
