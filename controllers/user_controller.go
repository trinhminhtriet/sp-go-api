package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/trinhminhtriet/sp-go-api/models"
)

type UserController struct {
	Controller
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

func (ctl UserController) GetAll(c *gin.Context) {
	users := []models.User{}
	ctl.db.Find(&users)

	ctl.SuccessResponse(c, gin.H{
		"users": users,
	})
}

func (ctl UserController) Get(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	ctl.db.First(&user, id)

	ctl.SuccessResponse(c, gin.H{
		"user": user,
	})
}

type userJSON struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

func (ctl UserController) Create(c *gin.Context) {
	var json userJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	user := models.User{
		Username: json.Username,
	}
	ctl.db.Where(&user).First(&user)

	if user.ID > 0 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "EXIST")
		return
	}

	ctl.db.Create(&user)
	ctl.SuccessResponse(c, gin.H{
		"user": user,
	})
}

func (ctl UserController) Update(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	ctl.db.First(&user, id)
	if user.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	var json userJSON
	if c.BindJSON(&json) != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.db.Model(&user).Updates(&models.User{
		Username: json.Username,
		Email:    json.Email,
		Password: json.Password,
		Status:   json.Status,
	})

	ctl.SuccessResponse(c, gin.H{
		"user": user,
	})
}

func (ctl UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}

	ctl.db.First(&user, id)
	if user.ID < 1 {
		ctl.ErrorResponse(c, http.StatusBadRequest, "INVALID")
		return
	}

	ctl.db.Delete(&user)
	ctl.SuccessResponse(c, gin.H{
		"user": user,
	})
}
