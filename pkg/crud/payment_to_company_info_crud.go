package crud

import (
	"github/Kotaro-Yamagishi/gorm-demo/pkg/models"

	"gorm.io/gorm"
)

type PaymentToCompanyInfoCRUD struct {
    db *gorm.DB
}

func NewPaymentToCompanyInfoCRUD(db *gorm.DB) *PaymentToCompanyInfoCRUD {
    return &PaymentToCompanyInfoCRUD{db}
}

func (crud *PaymentToCompanyInfoCRUD) Create(payment *models.PaymentToCompanyInfo) error {
    return crud.db.Create(payment).Error
}

func (crud *PaymentToCompanyInfoCRUD) GetByID(id string) (*models.PaymentToCompanyInfo, error) {
    var payment models.PaymentToCompanyInfo
    if err := crud.db.First(&payment, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &payment, nil
}

func (crud *PaymentToCompanyInfoCRUD) Update(payment *models.PaymentToCompanyInfo) error {
    return crud.db.Save(payment).Error
}

func (crud *PaymentToCompanyInfoCRUD) Delete(id string) error {
    return crud.db.Delete(&models.PaymentToCompanyInfo{}, "id = ?", id).Error
}
