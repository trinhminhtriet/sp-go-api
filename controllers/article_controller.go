package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/trinhminhtriet/sp-go-api/models"
)

type ArticleController struct {
	Controller
	db *gorm.DB
}

func NewArticleController(db *gorm.DB) *ArticleController {
	return &ArticleController{db: db}
}

func (ctl ArticleController) GetAll(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	articles := []models.Article{}
	ctl.db.Model(&user).Related(&articles)

	ctl.SuccessResponse(c, gin.H{
		"articles": articles,
	})
}

func (ctl ArticleController) Get(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	id := c.Param("id")
	article := models.Article{}
	ctl.db.Model(&user).Related(&[]models.Article{}).First(&article, id)

	if article.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.SuccessResponse(c, gin.H{
		"article": article,
	})
}

type CreateArticleJSON struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

func (ctl ArticleController) Create(c *gin.Context) {
	var json CreateArticleJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	user := c.MustGet("user").(models.User)
	article := models.Article{
		UserID: user.ID,
		Title:  json.Title,
		Body:   json.Body,
	}
	ctl.db.Create(&article)

	ctl.SuccessResponse(c, gin.H{
		"article": article,
	})
}

type updateArticleJSON struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (ctl ArticleController) Update(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	id := c.Param("id")
	article := models.Article{}
	ctl.db.Model(&user).Related(&[]models.Article{}).First(&article, id)
	if article.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	var json updateArticleJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.db.Model(&article).Updates(&models.Article{
		Title: json.Title,
		Body:  json.Body,
	})

	ctl.SuccessResponse(c, gin.H{
		"article": article,
	})
}

func (ctl ArticleController) Delete(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	id := c.Param("id")
	article := models.Article{}
	ctl.db.Model(&user).Related(&[]models.Article{}).First(&article, id)
	if article.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.db.Delete(&article)
	ctl.SuccessResponse(c, gin.H{
		"article": article,
	})
}
