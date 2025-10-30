package migrations

import (
	"github.com/RajaSunrise/pajakku/config"
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/models"
	"github.com/sirupsen/logrus"
)

func Migrate() {
	logrus.Info("Starting database migration")
	cfg := config.AppConfig
	err := databases.DB.AutoMigrate(
		&models.User{},
		&models.UserAuth{},
		&models.Role{},
		&models.PasswordResetToken{},
		&models.Attachment{},
		&models.AuditLog{},
		&models.TaxReturn{},
		&models.Invoice{},
		&models.Notification{},
		&models.Payment{},
		&models.TaxType{},
	)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to migrate database")
	}

	// Create default admin role if not exists
	var adminRole models.Role

	if err := databases.DB.Where("nama_role = ?", "admin").First(&adminRole).Error; err != nil {
		adminRole = models.Role{
			NamaRole:    "admin",
			Permissions: `{"all": true}`,
		}
		if err := databases.DB.Create(&adminRole).Error; err != nil {
			logrus.WithError(err).Fatal("Failed to create admin role")
		}
		logrus.Info("Admin role created")
	}

	var usersRole models.Role

	if err := databases.DB.Where("nama_role = ?", "users").First(&usersRole).Error; err != nil {
		usersRole = models.Role{
			NamaRole:    "users",
			Permissions: `{"all": false}`,
		}
		if err := databases.DB.Create(&usersRole).Error; err != nil {
			logrus.WithError(err).Fatal("Failed to create users role")
		}
		logrus.Info("Users role created")
	}

	// Create default admin user if not exists
	var adminUser models.User
	if err := databases.DB.Where("email = ?", cfg.Admin.Email).First(&adminUser).Error; err != nil {
		adminUser = models.User{
			NIK:             123456789,          // dummy NIK
			NPWP:            "1234567890123456", // dummy NPWP
			Nama:            "Admin",
			Email:           cfg.Admin.Email,
			Alamat:          "Admin Address",
			JenisWajibPajak: "individu",
			RoleID:          adminRole.ID,
		}
		if err := databases.DB.Create(&adminUser).Error; err != nil {
			logrus.WithError(err).Fatal("Failed to create admin user")
		}

		// Create UserAuth for admin
		adminAuth := models.UserAuth{
			UserID: adminUser.ID,
		}
		if err := adminAuth.HashPassword(cfg.Admin.Password); err != nil {
			logrus.WithError(err).Fatal("Failed to hash admin password")
		}
		if err := databases.DB.Create(&adminAuth).Error; err != nil {
			logrus.WithError(err).Fatal("Failed to create admin auth")
		}
		logrus.Info("Admin user created")
	}

	logrus.Info("Database migration completed successfully")
}
