package domain

import "time"

// Struktur Konsumen untuk menyimpan data konsumen
type Konsumen struct {
	ID           uint   `gorm:"primaryKey"`
	NIK          string `gorm:"uniqueIndex;not null"`
	FullName     string `gorm:"not null"`
	LegalName    string
	TempatLahir  string
	TanggalLahir string
	Gaji         float64
	FotoKTP      string
	FotoSelfie   string
	LoanLimit    float64 `gorm:"default:0"` // Total limit pinjaman konsumen
	Limit1       float64 `gorm:"default:0"` // Limit untuk tenor 1 bulan
	Limit2       float64 `gorm:"default:0"` // Limit untuk tenor 2 bulan
	Limit3       float64 `gorm:"default:0"` // Limit untuk tenor 3 bulan
	Limit6       float64 `gorm:"default:0"` // Limit untuk tenor 6 bulan
}

// Struktur Loan untuk menyimpan data pengajuan pinjaman
type Loan struct {
	ID             uint      `gorm:"primaryKey"`
	KonsumenID     uint      `gorm:"not null"`
	Amount         float64   `gorm:"not null"`
	Limit          float64   `gorm:"not null"`
	Status         string    `gorm:"default:'pending'"`
	SubmissionDate time.Time `gorm:"not null"`
}

// Struktur User untuk menyimpan data pengguna
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

// Struktur UserLogin untuk menangani data input saat login
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
