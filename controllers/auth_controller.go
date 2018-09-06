package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/trinhminhtriet/sp-go-api/models"
	"github.com/trinhminhtriet/sp-go-api/shared/jwtauth"
	"github.com/trinhminhtriet/sp-go-api/shared/passhash"
)

type AuthController struct {
	db *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{db}
}

type AuthJSON struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (controller AuthController) Auth(c *gin.Context) {
	var json AuthJSON
	if c.BindJSON(&json) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "パラメータが無効です",
		})
		return
	}

	var user models.User
	controller.db.Where(&models.User{
		Email: json.Email,
	}).First(&user)

	if user.ID < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "User not found",
		})
		return
	}

	err := passhash.ComparePassword(user.Password, json.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "INVALID_PASSWORD",
		})
		return
	}

	claims, err := jwtauth.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "トークンの生成に失敗しました",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "ok",
		"user":        user,
		"accessToken": claims,
	})
}
