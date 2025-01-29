package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yuskayus/pt-xyz-multifinance/internal/auth"
	"github.com/yuskayus/pt-xyz-multifinance/internal/domain"
	"golang.org/x/crypto/bcrypt" // Untuk bcrypt
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login meng-handle login dan menghasilkan JWT jika login berhasil
func (h *AuthHandler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user domain.User
	err := h.DB.Where("username = ?", input.Username).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Verifikasi password dengan bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT setelah login sukses
	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	log.Printf("Generated JWT: %s", token)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

// Middleware untuk melindungi endpoint dengan memverifikasi token JWT
func ParseJWT(c *gin.Context) (uint, error) {
	// Mengambil token dari header Authorization
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
		return 0, nil
	}

	// Menghapus prefix "Bearer" jika ada
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Verifikasi token
	userID, err := auth.ParseJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return 0, err
	}

	return userID, nil
}

// Endpoint yang membutuhkan autentikasi
func ProtectedEndpoint(c *gin.Context) {
	userID, err := ParseJWT(c)
	if err != nil {
		return
	}

	// Akses ke data yang hanya dapat diakses oleh pengguna yang sudah login
	c.JSON(http.StatusOK, gin.H{"user_id": userID, "message": "Welcome to the protected route!"})
}
