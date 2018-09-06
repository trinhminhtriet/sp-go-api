package controllers

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/trinhminhtriet/gorm-pagination/pagination"
	"github.com/trinhminhtriet/sp-go-api/models"
)

type SkillController struct {
	Controller
	db *gorm.DB
}

func NewSkillController(db *gorm.DB) *SkillController {
	return &SkillController{db: db}
}

func (ctl SkillController) GetAll(c *gin.Context) {
	skills := []models.Skill{}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10

	ctl.db.Find(&skills)

	result := pagination.Pagging(&pagination.Param{
		DB:      ctl.db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &skills)

	ctl.SuccessResponse(c, gin.H{
		"skill": result,
	})
}

func (ctl SkillController) Get(c *gin.Context) {
	id := c.Param("id")
	skill := models.Skill{}
	ctl.db.First(&skill, id)

	if skill.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.SuccessResponse(c, gin.H{
		"skill": skill,
	})
}

type skillJSON struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

func (ctl SkillController) Create(c *gin.Context) {
	var json skillJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	skill := models.Skill{
		Name: json.Name,
	}
	ctl.db.Where(&skill).First(&skill)

	if skill.ID > 0 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "EXIST")
		return
	}

	ctl.db.Create(&skill)
	ctl.SuccessResponse(c, gin.H{
		"skill": skill,
	})
}

func (ctl SkillController) Update(c *gin.Context) {
	id := c.Param("id")
	skill := models.Skill{}
	ctl.db.First(&skill, id)
	if skill.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	var json skillJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.db.Model(&skill).Updates(&models.Skill{
		Name:        json.Name,
		Description: json.Description,
		Status:      json.Status,
	})

	ctl.SuccessResponse(c, gin.H{
		"skill": skill,
	})
}

func (ctl SkillController) Delete(c *gin.Context) {
	id := c.Param("id")
	skill := models.Skill{}

	ctl.db.First(&skill, id)
	if skill.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.db.Delete(&skill)
	ctl.SuccessResponse(c, gin.H{
		"skill": skill,
	})
}
