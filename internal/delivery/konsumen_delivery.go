// internal/delivery/konsumen_handler.go
package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuskayus/pt-xyz-multifinance/internal/domain"
	"github.com/yuskayus/pt-xyz-multifinance/internal/service"
)

type KonsumenHandler struct {
	Service *service.KonsumenService
}

type TransactionRequest struct {
	KonsumenID uint    `json:"konsumen_id"`
	Tenor      int     `json:"tenor"`
	Amount     float64 `json:"amount"`
}

func (h *KonsumenHandler) ProcessTransaction(c *gin.Context) {
	var req TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.Service.ProcessTransaction(req.KonsumenID, req.Tenor, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaksi berhasil"})
}

// GetAll - Mendapatkan semua data konsumen
func (h *KonsumenHandler) GetAll(c *gin.Context) {
	konsumens, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get konsumens"})
		return
	}
	c.JSON(http.StatusOK, konsumens)
}

// Create - Membuat data konsumen baru
func (h *KonsumenHandler) Create(c *gin.Context) {
	var input domain.Konsumen
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validasi tambahan
	if input.NIK == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NIK is required"})
		return
	}

	err := h.Service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create konsumen"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Konsumen created successfully"})
}
