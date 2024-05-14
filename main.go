package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulgubili3003/jwt-user-authentication-golang/controllers"
	"github.com/rahulgubili3003/jwt-user-authentication-golang/db"
	initialize "github.com/rahulgubili3003/jwt-user-authentication-golang/init"
)

func init() {
	initialize.InitializeEnvVariables()
	db.ConnectToDb()
	db.SyncDB()
}

func main() {
	r := gin.Default()
	r.POST("/signUp", controllers.Signup)
	err := r.Run()
	if err != nil {
		return
	}
}
