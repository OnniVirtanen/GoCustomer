package auth

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var (
	secretKey []byte
)

func init() {
	errLoad := godotenv.Load()
	if errLoad != nil {
		log.Fatal("Error loading .env file at tokenService")
	}

	hexKey := os.Getenv("JWT_SECRET_KEY")
	if hexKey == "" {
		log.Fatal("JWT_SECRET_KEY environment variable is not set.")
	}

	var err error
	secretKey, err = hex.DecodeString(hexKey)
	if err != nil {
		log.Fatalf("Failed to decode JWT_SECRET_KEY: %v", err)
	}
}

func CreateToken(userId int, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"email":  email,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("Invalid token")
	}
	return nil
}
