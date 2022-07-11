package main

import (
	"contractor_panel/cmd"
	"os"
)

func main() {
	os.Setenv("PORT", "8080")
	os.Setenv("DATASOURCES_POSTGRES_HOST", "ec2-18-235-154-252.compute-1.amazonaws.com")
	os.Setenv("HEALTHCHECK_TIMEOUT", "30s")
	os.Setenv("DATASOURCES_POSTGRES_USER", "obktylemltwdzc")
	os.Setenv("DATASOURCES_POSTGRES_PASSWORD", "3492ecaa68258e934a6b71d0fa502f99d45ce6c159a588a403dbdd789aeedf84")
	os.Setenv("DATASOURCES_POSTGRES_DATABASE", "ddgndiu11pio51")
	os.Setenv("DATASOURCES_POSTGRES_PORT", "5432")
	os.Setenv("DATASOURCES_POSTGRES_SCHEMA", "public")
	os.Setenv("APP_NAME", "contractor-panel")
	os.Setenv("APP_INSTANCE", "instance1")
	os.Setenv("LOG_PRETTY_PRINT", "1")
	os.Setenv("CORS_ALLOWED_ORIGINS", "*")
	os.Setenv("CORS_ALLOWED_METHODS", "GET POST PUT DELETE OPTIONS")
	os.Setenv("CORS_ALLOWED_HEADERS", "Accept Authorization Content-Type X-CSRF-Token")
	os.Setenv("ACCESS_SECRET", "asdfertybvndfghrtyudfghasdfwer")
	os.Setenv("REFRESH_SECRET", "lhjfghjtyuiycvbnadfrtysdfzcvfg")
	//os.Setenv("REDIS_DSN", "rediss://:AVNS_BwpVUI_Q4IzmOVS@db-redis-blr1-07067-do-user-10184514-0.b.db.ondigitalocean.com:6379")
	os.Setenv("REDIS_DSN", "rediss://:AVNS_BwpVUI_Q4IzmOVS@db-redis-blr1-07067-do-user-10184514-0.b.db.ondigitalocean.com:25061")
	cmd.Execute()

}
