package router

import (
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/controllers"
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var 
(
	db *gorm.DB = database.SetupDbConnection()
	authController  controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	authRoutes := router.Group("auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	// userRoutes := router.Group("user", middlewares.AuthorizeJWT(jwtService))
	// {
	// 	userRoutes.GET("/profile", userController.Profile)
	// 	userRoutes.PUT("/profile", userController.Update)
	// }

	// photoRoutes := router.Group("photos", middlewares.AuthorizeJWT(jwtService))
	// {
	// 	photoRoutes.GET("/", photoController.All)
	// 	photoRoutes.POST("/", photoController.Insert)
	// 	photoRoutes.GET("/:id", photoController.FindByID)
	// 	photoRoutes.PUT("/:id", photoController.Update)
	// 	photoRoutes.DELETE("/:id", photoController.Delete)
	// }

	return router
}
