package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/trinhminhtriet/sp-go-api/models"
)

type JobController struct {
	Controller
	db *gorm.DB
}

func NewJobController(db *gorm.DB) *JobController {
	return &JobController{db: db}
}

func (ctl JobController) GetAll(c *gin.Context) {
	job := []models.Job{}
	ctl.db.Find(&job)

	ctl.SuccessResponse(c, gin.H{
		"jobs": job,
	})
}

func (ctl JobController) Get(c *gin.Context) {
	id := c.Param("id")
	job := models.Job{}
	ctl.db.First(&job, id)

	if job.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.SuccessResponse(c, gin.H{
		"job": job,
	})
}

type jobJSON struct {
	Title       string `json:"name" binding:"required"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func (ctl JobController) Create(c *gin.Context) {
	var json jobJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	job := models.Job{
		Title: json.Title,
	}
	ctl.db.Where(&job).First(&job)

	if job.ID > 0 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "EXIST")
		return
	}

	ctl.db.Create(&job)
	ctl.SuccessResponse(c, gin.H{
		"job": job,
	})
}

func (ctl JobController) Update(c *gin.Context) {
	id := c.Param("id")
	job := models.Job{}
	ctl.db.First(&job, id)
	if job.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	var json jobJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.db.Model(&job).Updates(&models.Job{
		Title:       json.Title,
		Description: json.Description,
		Status:      json.Status,
	})

	ctl.SuccessResponse(c, gin.H{
		"job": job,
	})
}

func (ctl JobController) Delete(c *gin.Context) {
	id := c.Param("id")
	job := models.Job{}

	ctl.db.First(&job, id)
	if job.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.db.Delete(&job)
	ctl.SuccessResponse(c, gin.H{
		"job": job,
	})
}
