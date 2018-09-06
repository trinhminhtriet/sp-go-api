package configs

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/trinhminhtriet/sp-go-api/controllers"
	"github.com/trinhminhtriet/sp-go-api/middleware"
)

func makeResource(r gin.IRouter, ctl controllers.ResourceController) {
	r.GET("", ctl.GetAll)
	r.GET("/:id", ctl.Get)
	r.POST("", ctl.Create)
	r.PUT("/:id", ctl.Update)
	r.PATCH("/:id", ctl.Update)
	r.DELETE("/:id", ctl.Delete)
}

func BuildRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Auth
	authController := controllers.NewAuthController(db)
	auth := router.Group("/auth")
	{
		auth.POST("/", authController.Auth)
	}

	unAuthorized := router.Group("/api/v1")
	{
		//Company
		userController := controllers.NewUserController(db)
		makeResource(unAuthorized.Group("/user"), userController)

		//Company
		companyController := controllers.NewCompanyController(db)
		makeResource(unAuthorized.Group("/company"), companyController)

		//Job
		jobController := controllers.NewJobController(db)
		makeResource(unAuthorized.Group("/job"), jobController)

		//Skill
		skillController := controllers.NewSkillController(db)
		makeResource(unAuthorized.Group("/skill"), skillController)

	}

	// Authentication required
	authorized := router.Group("/")
	authorized.Use(middleware.JWTMiddleware(db))
	{
		// Users
		userController := controllers.NewUserController(db)
		users := authorized.Group("/user")
		{
			users.GET("/", userController.GetAll)
			users.GET("/:id", userController.Get)
		}

	}

	return router
}