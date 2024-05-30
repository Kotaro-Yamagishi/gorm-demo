package models

import "time"

type PaymentToCompanyInfo struct {
    ID        string `gorm:"primaryKey"`
    SiteID    string
    Site      Site `gorm:"foreignKey:SiteID"`
    Amount    int
    PayedAt   time.Time
    CreatedAt time.Time
    UpdatedAt time.Time
}
