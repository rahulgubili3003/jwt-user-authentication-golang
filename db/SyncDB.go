package db

import "github.com/rahulgubili3003/jwt-user-authentication-golang/models"

func SyncDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
