package controllers

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/trinhminhtriet/gorm-pagination/pagination"
	"github.com/trinhminhtriet/sp-go-api/models"
)

type IndustryController struct {
	Controller
	db *gorm.DB
}

func NewIndustryController(db *gorm.DB) *IndustryController {
	return &IndustryController{db: db}
}

func (ctl IndustryController) GetAll(c *gin.Context) {
	industrys := []models.Industry{}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10

	ctl.db.Find(&industrys)

	result := pagination.Pagging(&pagination.Param{
		DB:      ctl.db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &industrys)

	ctl.SuccessResponse(c, gin.H{
		"industry": result,
	})
}

func (ctl IndustryController) Get(c *gin.Context) {
	id := c.Param("id")
	industry := models.Industry{}
	ctl.db.First(&industry, id)

	if industry.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.SuccessResponse(c, gin.H{
		"industry": industry,
	})
}

type industryJSON struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func (ctl IndustryController) Create(c *gin.Context) {
	var json industryJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	industry := models.Industry{
		Name: json.Name,
	}
	ctl.db.Where(&industry).First(&industry)

	if industry.ID > 0 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "EXIST")
		return
	}

	ctl.db.Create(&industry)
	ctl.SuccessResponse(c, gin.H{
		"industry": industry,
	})
}

func (ctl IndustryController) Update(c *gin.Context) {
	id := c.Param("id")
	industry := models.Industry{}
	ctl.db.First(&industry, id)
	if industry.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	var json industryJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.db.Model(&industry).Updates(&models.Industry{
		Name:        json.Name,
		Description: json.Description,
		Status:      json.Status,
	})

	ctl.SuccessResponse(c, gin.H{
		"industry": industry,
	})
}

func (ctl IndustryController) Delete(c *gin.Context) {
	id := c.Param("id")
	industry := models.Industry{}

	ctl.db.First(&industry, id)
	if industry.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.db.Delete(&industry)
	ctl.SuccessResponse(c, gin.H{
		"industry": industry,
	})
}
