package main

import (
	_ "github/Kotaro-Yamagishi/gorm-demo/docs"
	"github/Kotaro-Yamagishi/gorm-demo/pkg/handlers"
	"github/Kotaro-Yamagishi/gorm-demo/pkg/models"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title GORM Migrate Project API
// @version 1.0
// @description This is a sample server for managing sites and payment information.
// @host localhost:8080
// @BasePath /

func main() {
    dsn := "user:password@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to the database")
    }

    // Run migrations
    runMigrations(db)

    // Insert sample data
    insertSampleData(db)

    // Initialize Gin
    r := gin.Default()

    // Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Handlers
    siteHandler := handlers.NewSiteHandler(db)
    paymentHandler := handlers.NewPaymentToCompanyInfoHandler(db)

    // Routes
    siteGroup := r.Group("/sites")
    {
        siteGroup.POST("/", siteHandler.CreateSite)
        siteGroup.GET("/:id", siteHandler.GetSite)
        siteGroup.PUT("/:id", siteHandler.UpdateSite)
        siteGroup.DELETE("/:id", siteHandler.DeleteSite)
    }

    paymentGroup := r.Group("/payments")
    {
        paymentGroup.POST("/", paymentHandler.CreatePaymentToCompanyInfo)
        paymentGroup.GET("/:id", paymentHandler.GetPaymentToCompanyInfo)
        paymentGroup.PUT("/:id", paymentHandler.UpdatePaymentToCompanyInfo)
        paymentGroup.DELETE("/:id", paymentHandler.DeletePaymentToCompanyInfo)
    }

    r.Run(":8080")
}

func runMigrations(db *gorm.DB) {
    db.AutoMigrate(&models.Site{}, &models.PaymentToCompanyInfo{})
}

func insertSampleData(db *gorm.DB) {
    // Create a new site
    site := models.Site{
        ID:        "site1",
        Name:      "Example Site",
        Url:       "http://example.com",
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    db.Create(&site)

    // Create new payment info associated with the site
    paymentInfo1 := models.PaymentToCompanyInfo{
        ID:        "payment1",
        SiteID:    site.ID,
        Amount:    1000,
        PayedAt:   time.Now(),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    paymentInfo2 := models.PaymentToCompanyInfo{
        ID:        "payment2",
        SiteID:    site.ID,
        Amount:    2000,
        PayedAt:   time.Now(),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    db.Create(&paymentInfo1)
    db.Create(&paymentInfo2)
}
