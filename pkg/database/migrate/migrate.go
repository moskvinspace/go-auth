package migrate

import (
	"github.com/moskvinspace/simple-web-app/pkg/database"
	"github.com/moskvinspace/simple-web-app/pkg/models"
)

func MigrationDB() error {
	if err := database.DB.AutoMigrate(
		&models.User{},
	); err != nil {
		return err
	}

	return nil
}
