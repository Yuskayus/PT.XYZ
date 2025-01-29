package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yuskayus/pt-xyz-multifinance/config"
	"github.com/yuskayus/pt-xyz-multifinance/internal/auth"
	"github.com/yuskayus/pt-xyz-multifinance/internal/delivery"
	"github.com/yuskayus/pt-xyz-multifinance/internal/domain"
	"github.com/yuskayus/pt-xyz-multifinance/internal/handler"
	"github.com/yuskayus/pt-xyz-multifinance/internal/repository"
	"github.com/yuskayus/pt-xyz-multifinance/internal/service"
)

// Middleware untuk melindungi endpoint dengan memverifikasi token JWT
func ParseJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mengambil token dari header Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			c.Abort() // Menghentikan eksekusi middleware lebih lanjut
			return
		}

		// Menghapus prefix "Bearer" jika ada
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Verifikasi token
		userID, err := auth.ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort() // Menghentikan eksekusi middleware lebih lanjut
			return
		}

		// Menyimpan userID ke context untuk digunakan di handler selanjutnya
		c.Set("userID", userID)

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}

func main() {
	// Inisialisasi database
	db := config.InitDB()

	// AutoMigrate untuk membuat tabel jika belum ada
	if err := db.AutoMigrate(&domain.Konsumen{}, &domain.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Dependency injection untuk Konsumen
	konsumenRepo := &repository.KonsumenRepository{DB: db}
	konsumenService := &service.KonsumenService{Repo: konsumenRepo}
	konsumenHandler := &delivery.KonsumenHandler{Service: konsumenService}

	// Dependency injection untuk AuthHandler
	authHandler := &handler.AuthHandler{DB: db}

	// Gin router
	r := gin.Default()

	// Route setup
	r.GET("/konsumen", konsumenHandler.GetAll)
	r.POST("/konsumen", konsumenHandler.Create)

	// Endpoint untuk login
	r.POST("/login", authHandler.Login)

	// Protected route yang memerlukan autentikasi
	protected := r.Group("/protected")
	protected.Use(ParseJWTMiddleware()) // Menggunakan middleware untuk verifikasi JWT
	{
		protected.GET("", konsumenHandler.GetAll) // Akses ke data yang dilindungi
	}

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
