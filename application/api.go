package application

import (
	"context"
	"contractor_panel/application/cerrors"
	"contractor_panel/application/config"
	"contractor_panel/application/controller"
	"contractor_panel/application/cvalidator"
	"contractor_panel/application/middleware"
	"contractor_panel/application/respond"
	"contractor_panel/application/service"
	"contractor_panel/infrastructure/logging"
	"contractor_panel/infrastructure/persistence/postgres"
	"contractor_panel/infrastructure/persistence/redis"
	"github.com/etherlabsio/healthcheck"
	redis2 "github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"net/http"
)

// NewApi конфигурирует API
func NewApi() (http.Handler, error) {
	logging.ConfigureLogger()
	cvalidator.ConfigureValidator()

	pc := postgres.DBConn()
	redisClient := redis.RedisConn()

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(handleNotFoundError)

	configureHealthchecks(r, pc)

	r.Use(middleware.CorrelationHandler)
	r.Use(middleware.RequestLoggerHandler)
	r.Use(middleware.RecoveryHandler(middleware.PrintRecoveryStack(true)))

	r.Use(middleware.CORS(
		middleware.AllowedOrigins(viper.GetStringSlice(config.CorsAllowedOrigins)),
		middleware.AllowedMethods(viper.GetStringSlice(config.CorsAllowedMethods)),
		middleware.AllowedHeaders(viper.GetStringSlice(config.CorsAllowedHeaders)),
		middleware.AllowCredentials()))

	err := configureRoutes(r, pc, redisClient)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func configureHealthchecks(r *mux.Router, pc *pgxpool.Pool) {
	r.Handle("/health", healthcheck.Handler(
		healthcheck.WithTimeout(viper.GetDuration(config.HealthcheckTimeout)),

		healthcheck.WithChecker("database", healthcheck.CheckerFunc(
			func(ctx context.Context) error {
				return pc.Ping(ctx)
			}),
		),
	))
}

func configureRoutes(r *mux.Router, pc *pgxpool.Pool, client *redis2.Client) error {
	signRepo := postgres.NewSignRepository(pc)
	userRepo := postgres.NewUserRepository(pc)
	tokenRepo := redis.NewTokenRepository(client)
	contractTemplateRepo := postgres.NewContractTemplateRepository(pc)
	//rbReport := postgres.NewReportTemplateRepository(pc)

	signService := service.NewSignService(signRepo, tokenRepo)
	contractTemplateService := service.NewContractTemplateService(contractTemplateRepo)

	//region Sign routes
	apiSign := r.PathPrefix("/api/v1/sign").Subrouter()
	controller.NewSignController(signService).HandleRoutes(apiSign)
	//end region

	//region Contractor routes
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.AuthHandler(userRepo))
	controller.NewContractTemplateController(contractTemplateService).HandleRoutes(api)

	//endregion

	return nil
}

func handleNotFoundError(w http.ResponseWriter, r *http.Request) {
	respond.WithError(w, r, cerrors.ErrResourceNotFound(r))
}
