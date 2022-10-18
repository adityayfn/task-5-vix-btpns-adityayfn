package router

import (
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/controllers"
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/database"
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/middlewares"
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/repository"
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var 
(
	db              *gorm.DB                    = database.SetupDbConnection()
	userRepository  repository.UserRepository   = repository.NewUserRepository(db)
	photoRepository repository.PhotoRepository  = repository.NewPhotoRepository(db)
	jwtService      service.JWTService          = service.NewJWTService()
	userService     service.UserService         = service.NewUserService(userRepository)
	photoService    service.PhotoService        = service.NewPhotoService(photoRepository)
	authService     service.AuthService         = service.NewAuthService(userRepository)
	authController  controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
	userController  controllers.UserController  = controllers.NewUserController(userService, jwtService)
	photoController controllers.PhotoController = controllers.NewPhotoController(photoService, jwtService)
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	authRoutes := router.Group("auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	

	userRoutes := router.Group("users",middlewares.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile/:id", userController.Update)
		userRoutes.DELETE("/:id",userController.Delete)
	}
	

	photoRoutes := router.Group("photos", middlewares.AuthorizeJWT(jwtService))
	{
		photoRoutes.GET("/", photoController.All)
		photoRoutes.POST("/", photoController.Insert)
		photoRoutes.GET("/:id", photoController.FindByID)
		photoRoutes.PUT("/:id", photoController.Update)
		photoRoutes.DELETE("/:id", photoController.Delete)
	}

	return router
}
