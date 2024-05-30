package models

import "time"

type Site struct {
    ID          string `gorm:"primaryKey"`
    Name        string
    Url         string
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeleteFlag  bool
    Payments    []PaymentToCompanyInfo `gorm:"foreignKey:SiteID"`
}
