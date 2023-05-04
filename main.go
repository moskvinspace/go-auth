package main

import (
	"github.com/moskvinspace/simple-web-app/pkg/api"
	"github.com/moskvinspace/simple-web-app/pkg/database"
	"github.com/moskvinspace/simple-web-app/pkg/database/migrate"
)

func init() {
	database.ConnectDb()
	migrate.MigrationDB()
}

func main() {
	api.Start()
}
