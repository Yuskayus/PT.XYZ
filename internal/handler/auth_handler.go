package handler

import (
	"log"
	"net/http"
	"strings"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuskayus/pt-xyz-multifinance/internal/auth"
	"github.com/yuskayus/pt-xyz-multifinance/internal/domain"
	"golang.org/x/crypto/bcrypt" // Untuk bcrypt
	"gorm.io/gorm"
)

type LoanHandler struct {
	DB *gorm.DB
}

type LoanInput struct {
	KonsumenID uint    `json:"konsumen_id"`
	Amount     float64 `json:"amount"`
}

// Endpoint untuk mengajukan pinjaman
func (h *LoanHandler) ApplyLoan(c *gin.Context) {
	var input LoanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Cek apakah customer (Konsumen) ada
	var konsumen domain.Konsumen
	if err := h.DB.First(&konsumen, input.KonsumenID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	// Cek apakah jumlah pengajuan lebih besar dari limit pinjaman
	if input.Amount > konsumen.LoanLimit {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Loan amount exceeds the limit"})
		return
	}

	// Membuat pengajuan pinjaman
	loan := domain.Loan{
		KonsumenID:     input.KonsumenID,
		Amount:         input.Amount,
		Limit:          konsumen.LoanLimit,
		Status:         "pending", // Status awal adalah pending
		SubmissionDate: time.Now(),
	}

	if err := h.DB.Create(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to apply for loan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Loan application submitted successfully",
		"loan_id": loan.ID,
	})
}

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
