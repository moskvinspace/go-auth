package main

import (
	"github.com/moskvinspace/simple-web-app/pkg/api"
	"github.com/moskvinspace/simple-web-app/pkg/database"
	"github.com/moskvinspace/simple-web-app/pkg/database/migrate"
)

func init() {
	database.ConnectDB()
	migrate.MigrationDB()
}

func main() {
	api.Start()
}
