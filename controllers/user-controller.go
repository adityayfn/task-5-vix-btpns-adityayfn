package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/adityayfn/task-5-vix-btpns-adityayfn/app"
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/helpers"
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/models"
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
	Delete(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateModel models.UserUpdateModel
	errModel := context.ShouldBind(&userUpdateModel)
	if errModel != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errModel.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateModel.ID = id
	u := c.userService.Update(userUpdateModel)
	res := helpers.BuildResponse(true, "success", u)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helpers.BuildResponse(true, "success", user)
	context.JSON(http.StatusOK, res)

}

func(c *userController) Delete(context *gin.Context){
	var user app.User
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to get id", "No param id were found", helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}

	user.ID = id
	c.userService.Delete(user)
	res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObj{})
	context.JSON(http.StatusOK, res)
}