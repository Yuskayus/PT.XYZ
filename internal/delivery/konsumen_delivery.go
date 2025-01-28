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

func (h *KonsumenHandler) GetAll(c *gin.Context) {
	konsumens, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get konsumens"})
		return
	}
	c.JSON(http.StatusOK, konsumens)
}

func (h *KonsumenHandler) Create(c *gin.Context) {
	var input domain.Konsumen
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create konsumen"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Konsumen created successfully"})
}
