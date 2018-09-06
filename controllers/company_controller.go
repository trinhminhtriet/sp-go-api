package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/trinhminhtriet/sp-go-api/models"
)

type CompanyController struct {
	Controller
	db *gorm.DB
}

func NewCompanyController(db *gorm.DB) *CompanyController {
	return &CompanyController{db: db}
}

func (ctl CompanyController) GetAll(c *gin.Context) {
	company := []models.Company{}
	ctl.db.Find(&company)

	ctl.SuccessResponse(c, gin.H{
		"companies": company,
	})
}

func (ctl CompanyController) Get(c *gin.Context) {
	id := c.Param("id")
	company := models.Company{}
	ctl.db.Preload("Jobs").First(&company, id)

	if company.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.SuccessResponse(c, gin.H{
		"company": company,
	})
}

type companyJSON struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func (ctl CompanyController) Create(c *gin.Context) {
	var json companyJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	company := models.Company{
		Name: json.Name,
	}
	ctl.db.Where(&company).First(&company)

	if company.ID > 0 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "EXIST")
		return
	}

	ctl.db.Create(&company)
	ctl.SuccessResponse(c, gin.H{
		"company": company,
	})
}

func (ctl CompanyController) Update(c *gin.Context) {
	id := c.Param("id")
	company := models.Company{}
	ctl.db.First(&company, id)
	if company.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	var json companyJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.db.Model(&company).Updates(&models.Company{
		Name:        json.Name,
		Description: json.Description,
		Status:      json.Status,
	})

	ctl.SuccessResponse(c, gin.H{
		"company": company,
	})
}

func (ctl CompanyController) Delete(c *gin.Context) {
	id := c.Param("id")
	company := models.Company{}

	ctl.db.First(&company, id)
	if company.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.db.Delete(&company)
	ctl.SuccessResponse(c, gin.H{
		"company": company,
	})
}
