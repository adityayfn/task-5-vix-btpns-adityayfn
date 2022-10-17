package controllers

import (
	"net/http"
	"strconv"

	"github.com/adityayfn/task-5-vix-btpns-adityayfn/app"
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/helpers"
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/models"
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/service"
	"github.com/gin-gonic/gin"
)


type AuthController interface{
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginModel models.LoginModel
	errModel := ctx.ShouldBind(&loginModel)
	if errModel != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errModel.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginModel.Email, loginModel.Password)
	if v, ok := authResult.(app.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10), v.Username, v.Email)
		v.Token = generatedToken
		response := helpers.BuildResponse(true, "success", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helpers.BuildErrorResponse("Please check again your credential", "Invalid Credential", helpers.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerModel models.RegisterModel
	errModel := ctx.ShouldBind(&registerModel)
	if errModel != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errModel.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerModel.Email) {
		response := helpers.BuildErrorResponse("Failed to process request", "Duplicate email", helpers.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerModel)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10), createdUser.Username, createdUser.Email)
		createdUser.Token = token
		response := helpers.BuildResponse(true, "success", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
