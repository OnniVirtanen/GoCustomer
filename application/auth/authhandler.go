package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	var authHandler = AuthHandler{}
	return &authHandler
}

func (h *AuthHandler) Login(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	var u User
	json.NewDecoder(c.Request.Body).Decode(&u)
	fmt.Printf("The user request value %v", u)

	if u.Username == "Chek" && u.Password == "123456" {
		tokenString, err := CreateToken(u.Username)
		if err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			//fmt.Errorf("No username found")
		}
		c.Writer.WriteHeader(http.StatusOK)
		fmt.Fprint(c.Writer, tokenString)
		return
	} else {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(c.Writer, "Invalid credentials")
	}
}
