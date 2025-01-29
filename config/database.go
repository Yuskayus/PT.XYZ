package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv" // Import godotenv untuk memuat file .env
	"github.com/yuskayus/pt-xyz-multifinance/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	// Memuat file .env jika ada
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using environment variables directly.")
	}

	// Mengambil konfigurasi dari variabel lingkungan
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal(`
DATABASE_URL environment variable is not set.
Please define it in your .env file or directly in the environment.

Example:
DATABASE_URL="host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
`)
	}

	// Membuka koneksi ke database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Mengaktifkan logging query jika diperlukan
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Menambahkan pengaturan pool koneksi jika dibutuhkan
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get generic database object: %v", err)
	}
	sqlDB.SetMaxOpenConns(10)                  // Set jumlah koneksi terbuka maksimal
	sqlDB.SetMaxIdleConns(5)                   // Set jumlah koneksi idle
	sqlDB.SetConnMaxLifetime(10 * time.Minute) // Set waktu maksimum koneksi

	// Lakukan migrasi untuk tabel konsumen dan lainnya
	err = db.AutoMigrate(&domain.Konsumen{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}

	log.Println("Database connection established and migration completed.")
	return db
}
