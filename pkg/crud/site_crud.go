package crud

import (
	"github/Kotaro-Yamagishi/gorm-demo/pkg/models"

	"gorm.io/gorm"
)

type SiteCRUD struct {
    db *gorm.DB
}

func NewSiteCRUD(db *gorm.DB) *SiteCRUD {
    return &SiteCRUD{db}
}

func (crud *SiteCRUD) Create(site *models.Site) error {
    return crud.db.Create(site).Error
}

func (crud *SiteCRUD) GetByID(id string) (*models.Site, error) {
    var site models.Site
    if err := crud.db.Preload("Payments").First(&site, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &site, nil
}

func (crud *SiteCRUD) Update(site *models.Site) error {
    return crud.db.Save(site).Error
}

func (crud *SiteCRUD) Delete(id string) error {
    return crud.db.Delete(&models.Site{}, "id = ?", id).Error
}
