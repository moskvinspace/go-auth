package migrate

import (
	"github.com/moskvinspace/go-auth/pkg/database"
	"github.com/moskvinspace/go-auth/pkg/models"
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
