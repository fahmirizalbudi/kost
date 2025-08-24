package main

import (
	"api/configs"
	"api/database/migrations"
	"api/helpers"
	"api/router"

	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	helpers.LoadENV()
	configs.GetPostgresConnection()
	migrations.Run(configs.DB, migrate.Up)
	defer configs.DB.Close()

	router.Setup().Run()
}