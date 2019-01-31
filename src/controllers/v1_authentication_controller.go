package controllers

import (
	"../constants"
	"../helpers"
	"../middleware"
	"../objects"
	"../services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type V1AuthenticationController struct {
	service     services.V1AuthenticationService
	errorHelper helpers.ErrorHelper
}

func V1AuthenticationControllerHandler(router *gin.Engine) {

	handler := &V1AuthenticationController{
		errorHelper: helpers.ErrorHelperHandler(),
		service:     services.V1AuthenticationServiceHandler(),
	}

	defaultMiddleware := middleware.DefaultMiddleware{}

	group := router.Group("v1/authentication")
	group.Use(defaultMiddleware.AuthenticationMiddleware())
	{
		group.GET("profile", handler.GetProfile)
		group.POST("generate", handler.Generate)
	}

}

func (handler *V1AuthenticationController) GetProfile(context *gin.Context) {

	currentUser, exists := context.Get("user")

	if exists {
		context.JSON(http.StatusOK, currentUser)
	} else {
		context.JSON(http.StatusNoContent, nil)
	}

}

func (handler *V1AuthenticationController) Generate(context *gin.Context) {

	requestObject := objects.V1AuthenticationObjectRequest{}
	context.ShouldBind(&requestObject)

	result, err := handler.service.Generate(requestObject)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)

}
