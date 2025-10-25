package migrations

import (
	"fmt"
	"log"
	"os"

	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/models"
)

func Migrate() {
	err := databases.DB.AutoMigrate(
		&models.UserProfile{},
		&models.UserAuth{},
		&models.Billing{},
		&models.ReportSPT{},
		&models.PasswordResetToken{},
	)
	if err != nil {
		panic(err)
	}
	if _, err := fmt.Fprintln(os.Stdout, []any{"Succes to migrate"}); err != nil {
		log.Fatalf("Failed to print message: %v", err)
	}

}
