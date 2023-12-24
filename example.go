package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// path paramter with name details will mapped to Details
type URI struct {
	Details string `json:"name" uri:"details"`
}

func main() {
	engine := gin.New()
	// adding path params to router
	engine.GET("/test/:details", func(context *gin.Context) {
		uri := URI{}
		// binding to URI
		if err := context.BindUri(&uri); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(uri)
		context.JSON(http.StatusAccepted, &uri)
	})
	engine.Run(":3000")
}
