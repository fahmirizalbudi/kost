package configs

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"os"
)

var (
	Postgres *sql.DB
	err error
)

func GetPostgresConnection() {
	postgresInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	Postgres, err = sql.Open("postgres", postgresInfo)
	if err != nil {
		panic(err)
	}

	err = Postgres.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("PostgreSQL connection established successfully.")
}
