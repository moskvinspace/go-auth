package main

import (
	"github.com/moskvinspace/go-auth/api"
	"github.com/moskvinspace/go-auth/database"
	"github.com/moskvinspace/go-auth/database/migrate"
)

func init() {
	database.ConnectDB()
	migrate.MigrationDB()
}

func main() {
	api.Start()
}
