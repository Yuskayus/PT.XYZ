package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yuskayus/pt-xyz-multifinance/config"
	"github.com/yuskayus/pt-xyz-multifinance/internal/delivery"
	"github.com/yuskayus/pt-xyz-multifinance/internal/domain" // Import domain package
	"github.com/yuskayus/pt-xyz-multifinance/internal/repository"
	"github.com/yuskayus/pt-xyz-multifinance/internal/service"
)

func main() {
	// Inisialisasi database
	db := config.InitDB()

	// AutoMigrate untuk membuat tabel jika belum ada
	if err := db.AutoMigrate(&domain.Konsumen{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Dependency injection
	konsumenRepo := &repository.KonsumenRepository{DB: db}
	konsumenService := &service.KonsumenService{Repo: konsumenRepo}
	konsumenHandler := &delivery.KonsumenHandler{Service: konsumenService}

	// Gin router
	r := gin.Default()

	// Route setup
	r.GET("/konsumen", konsumenHandler.GetAll)
	r.POST("/konsumen", konsumenHandler.Create)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
