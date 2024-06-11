package main

import (
	"baseapi/config"
	"baseapi/db"
	"baseapi/handlers/auth/register"
	"baseapi/handlers/auth/login"
	"baseapi/handlers/auth/logout"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db.InitDB(cfg)

	r := gin.Default()
	
	r.POST("/baseapi/auth/register", register.Register)
	r.POST("/baseapi/auth/login", login.Login)
	r.POST("/baseapi/auth/logout", logout.Logout)

	r.Run(":8080")
}
