package main

import (
	"api/configs"
	"api/helpers"
)

func main() {
	helpers.LoadENV()
	configs.GetPostgresConnection()
	
	defer configs.Postgres.Close()
}