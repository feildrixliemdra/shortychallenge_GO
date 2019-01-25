package controllers

import (
	"../helpers"
	"../objects"
	"../services"
	"github.com/gin-gonic/gin"
	"github.com/ralali/agent_api/constants"
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

	group := router.Group("v1/authentication")
	{
		group.POST("generate", handler.Generate)
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
