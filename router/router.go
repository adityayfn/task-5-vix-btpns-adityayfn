package router

import (
	"github.com/gin-gonic/gin"
)

var 
{
	db *gorm.DB = config.SetupDbConnection()
	authController controllerAuthController = controller.NewAuthController()


}

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
