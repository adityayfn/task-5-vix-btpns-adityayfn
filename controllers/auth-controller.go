package controllers

import "github.com/gin-gonic/gin"


type AuthController interface{
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct{
// 
}

func NewAuthController() AuthController{
	return &authController{}
}
func(c *authController)Login(ctx *gin.Context){
	ctx.JSON(http.StatusOk,gin.H{
		"message" : "Hello Login",
	})
}

func(c *authController)Register(ctx *gin.Context){
	ctx.JSON(http.StatusOk,gin.H{
		"message" : "Hello Register",
	})
}