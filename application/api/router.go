package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, db *sql.DB) {
	SetupAuthRouter(router, db)
	SetupCustomerRouter(router, db)
	SetupProductRouter(router, db)
}
