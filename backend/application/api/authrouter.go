package api

import (
	"database/sql"

	"example.com/backend/application/auth"
	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(router *gin.Engine, db *sql.DB) {
	authHandler := auth.NewAuthHandler(db)

	auth := router.Group("v1/auth/user")
	{
		auth.POST("/login", authHandler.LoginUser)
		auth.POST("/register", authHandler.RegisterUser)
	}
}
