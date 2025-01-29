package domain

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
