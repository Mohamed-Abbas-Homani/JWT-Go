package initializers

import (
	"github.com/Mohamed-Abbas-Homani/jwt-go/models"
)

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}