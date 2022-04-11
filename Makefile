default: migrations-install-tool

dbstring := "postgres://${DATASOURCES_POSTGRES_USER}:${DATASOURCES_POSTGRES_PASSWORD}@${DATASOURCES_POSTGRES_HOST}:${DATASOURCES_POSTGRES_PORT}/${DATASOURCES_POSTGRES_DATABASE}?sslmode=disable"

# Устанавливает библиотеку для миграций.
migrations-install-tool:
	@go install github.com/pressly/goose/v3/cmd/goose@v3.5.0

# Устанавливает миграции.
migrations-up:
	@goose -dir=migrations postgres ${dbstring} up

# Статус миграций в ДБ.
migrations-status:
	@echo ${dbstring}
	@goose -dir=migrations postgres ${dbstring} status

.PHONY: default migrations-install-tool migrations-up migrations-status