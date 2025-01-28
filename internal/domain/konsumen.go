package domain

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
