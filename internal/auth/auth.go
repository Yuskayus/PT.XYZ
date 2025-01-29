package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Secret key untuk menandatangani JWT
var jwtKey = []byte("your-secret-key")

// Fungsi untuk membuat JWT
func GenerateJWT(userID uint) (string, error) {
	// Membuat klaim JWT
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // Token akan kedaluwarsa dalam 1 jam
	}

	// Membuat token dengan klaim yang sudah dibuat
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Menandatangani token dengan secret key
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error signing token:", err)
		return "", err
	}

	return signedToken, nil
}

// Fungsi untuk mengurai dan memverifikasi JWT
func ParseJWT(tokenString string) (uint, error) {
	// Parse token dengan secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return 0, err
	}

	// Verifikasi klaim dan ambil user_id dari klaim
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(float64); ok {
			return uint(userID), nil
		}
		return 0, fmt.Errorf("user_id claim is missing or invalid")
	}

	return 0, fmt.Errorf("invalid token")
}
