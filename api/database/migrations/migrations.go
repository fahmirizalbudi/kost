package migrations

import (
    "database/sql"
    "embed"
    "fmt"
    migrate "github.com/rubenv/sql-migrate"
)

//go:embed *.sql
var dbMigrations embed.FS

func Run(dbParam *sql.DB, direction migrate.MigrationDirection) {
    migrations := migrate.EmbedFileSystemMigrationSource{
       FileSystem: dbMigrations,
       Root:       ".",
    }

    n, errs := migrate.Exec(dbParam, "postgres", migrations, direction)
    if errs != nil {
       panic(errs)
    }

    fmt.Println("Migration success, applied", n, "migrations!")
}