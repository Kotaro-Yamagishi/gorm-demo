package handlers

import (
	"github/Kotaro-Yamagishi/gorm-demo/pkg/crud"
	"github/Kotaro-Yamagishi/gorm-demo/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaymentToCompanyInfoHandler struct {
    Crud *crud.PaymentToCompanyInfoCRUD
}

func NewPaymentToCompanyInfoHandler(db *gorm.DB) *PaymentToCompanyInfoHandler {
    return &PaymentToCompanyInfoHandler{
        Crud: crud.NewPaymentToCompanyInfoCRUD(db),
    }
}

func (h *PaymentToCompanyInfoHandler) CreatePaymentToCompanyInfo(c *gin.Context) {
    var payment models.PaymentToCompanyInfo
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Crud.Create(&payment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, payment)
}

func (h *PaymentToCompanyInfoHandler) GetPaymentToCompanyInfo(c *gin.Context) {
    id := c.Param("id")
    payment, err := h.Crud.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
        return
    }
    c.JSON(http.StatusOK, payment)
}

func (h *PaymentToCompanyInfoHandler) UpdatePaymentToCompanyInfo(c *gin.Context) {
    id := c.Param("id")
    var payment models.PaymentToCompanyInfo
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    payment.ID = id
    if err := h.Crud.Update(&payment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, payment)
}

func (h *PaymentToCompanyInfoHandler) DeletePaymentToCompanyInfo(c *gin.Context) {
    id := c.Param("id")
    if err := h.Crud.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Payment deleted"})
}
