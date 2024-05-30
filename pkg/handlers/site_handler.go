package handlers

import (
    "net/http"
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
    "github/Kotaro-Yamagishi/gorm-demo/pkg/models"
    "github/Kotaro-Yamagishi/gorm-demo/pkg/crud"
)

type SiteHandler struct {
    Crud *crud.SiteCRUD
}

func NewSiteHandler(db *gorm.DB) *SiteHandler {
    return &SiteHandler{
        Crud: crud.NewSiteCRUD(db),
    }
}

// CreateSite godoc
// @Summary Create a new site
// @Description Create a new site with the provided information
// @Tags sites
// @Accept json
// @Produce json
// @Param site body models.Site true "Site"
// @Success 201 {object} models.Site
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /sites [post]
func (h *SiteHandler) CreateSite(c *gin.Context) {
    var site models.Site
    if err := c.ShouldBindJSON(&site); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }
    if err := h.Crud.Create(&site); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusCreated, site)
}

// GetSite godoc
// @Summary Get site by ID
// @Description Get site details by ID
// @Tags sites
// @Produce json
// @Param id path string true "Site ID"
// @Success 200 {object} models.Site
// @Failure 404 {object} models.ErrorResponse
// @Router /sites/{id} [get]
func (h *SiteHandler) GetSite(c *gin.Context) {
    id := c.Param("id")
    site, err := h.Crud.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Site not found"})
        return
    }
    c.JSON(http.StatusOK, site)
}

// UpdateSite godoc
// @Summary Update site by ID
// @Description Update site details by ID
// @Tags sites
// @Accept json
// @Produce json
// @Param id path string true "Site ID"
// @Param site body models.Site true "Updated Site"
// @Success 200 {object} models.Site
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /sites/{id} [put]
func (h *SiteHandler) UpdateSite(c *gin.Context) {
    id := c.Param("id")
    var site models.Site
    if err := c.ShouldBindJSON(&site); err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
        return
    }
    site.ID = id
    if err := h.Crud.Update(&site); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, site)
}

// DeleteSite godoc
// @Summary Delete site by ID
// @Description Delete a site by ID
// @Tags sites
// @Produce json
// @Param id path string true "Site ID"
// @Success 200 {object} models.MessageResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /sites/{id} [delete]
func (h *SiteHandler) DeleteSite(c *gin.Context) {
    id := c.Param("id")
    if err := h.Crud.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, models.MessageResponse{Message: "Site deleted"})
}
