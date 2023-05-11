package migrate

import (
	"github.com/moskvinspace/go-auth/database"
	"github.com/moskvinspace/go-auth/models"
	"log"
	"os"
)

func MigrationDB() {
	if err := database.DB.AutoMigrate(
		&models.User{},
	); err != nil {
		log.Fatal("Failed to auto migrate. \n", err)
		os.Exit(2)
	}
}
