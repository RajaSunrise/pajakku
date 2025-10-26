package migrations

import (
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/models"
	"github.com/sirupsen/logrus"
)

func Migrate() {
	logrus.Info("Starting database migration")
	err := databases.DB.AutoMigrate(
		&models.UserProfile{},
		&models.UserAuth{},
		&models.Billing{},
		&models.ReportSPT{},
		&models.PasswordResetToken{},
	)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to migrate database")
	}
	logrus.Info("Database migration completed successfully")
}
