package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DB *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

func (h *AuthHandler) LoginUser(c *gin.Context) {
	userReqBody := new(UserRequestBody)
	if err := json.NewDecoder(c.Request.Body).Decode(userReqBody); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.Writer.Write([]byte("Please provide the correct input!"))
		return
	}
	query := `SELECT * FROM user where email = ?`
	rows, err := h.DB.Query(query, userReqBody.Email)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.Writer.Write([]byte("Please provide the correct input!"))
		return
	}
	var user *User
	for rows.Next() {
		user, err = ScanRow(rows)
		if err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if !checkPassword(user.Hash, userReqBody.Password) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.Writer.Write([]byte("Incorrect password please check again"))
		return
	}

	tokenString, err := CreateToken(user.Id, user.Email)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(tokenString))
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	userReqBody := new(UserRequestBody)
	if err := json.NewDecoder(c.Request.Body).Decode(userReqBody); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.Writer.Write([]byte("Please provide the correct input!!"))
		return
	}

	hashPassword, err := getHashPassword(userReqBody.Password)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	query := `INSERT INTO user (email, hash) VALUES (?, ?)`
	result, err := h.DB.Exec(query, userReqBody.Email, hashPassword)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	recordId, _ := result.LastInsertId()
	response := Response{
		Id: int(recordId),
	}

	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(response)
}

func getHashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
