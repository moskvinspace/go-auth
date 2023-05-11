package main

import (
	"github.com/moskvinspace/go-auth/pkg/api"
	"github.com/moskvinspace/go-auth/pkg/database"
	"github.com/moskvinspace/go-auth/pkg/database/migrate"
)

func init() {
	database.ConnectDB()
	migrate.MigrationDB()
}

func main() {
	api.Start()
}
