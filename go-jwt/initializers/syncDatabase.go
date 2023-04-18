package initializers

import "github.com/sushanpth/learn-go/go-jwt/models"

func SyncDatabase() {
	// we can only use DB instead initializers.DB with in initializers package
	DB.AutoMigrate(&models.User{})
}
