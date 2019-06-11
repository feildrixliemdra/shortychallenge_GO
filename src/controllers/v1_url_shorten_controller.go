package controllers

import(
	"../constants"
	"../helpers"
	// "../middleware"
	 "../objects"
	"../services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type V1UrlController struct {
	urlService services.V1UrlService
	errorHelper helpers.ErrorHelper
}

func V1UrlControllerHandler(router *gin.Engine) {

	handler:= &V1UrlController{
		urlService: services.V1UrlServiceHandler(),
		errorHelper: helpers.ErrorHelperHandler(),
}


	group:= router.Group("/shorten")

	group.GET(":url",handler.GetShortenUrl)
	group.GET(":url/stats", handler.GetUrlStats)
	group.POST("",handler.CreateNewUrl)

}

func (handler *V1UrlController) CreateNewUrl(context *gin.Context){

	requestObject := objects.V1NewUrlObjectRequest{}
	context.ShouldBind(&requestObject)

	result, err := handler.urlService.CreateNewUrl(requestObject.InputUrl, requestObject.ShortenUrl)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}	
	
	context.JSON(http.StatusCreated, gin.H{
		"status":"created successfully!",
		"data": result,
	})
}

func (handler *V1UrlController) GetShortenUrl(context *gin.Context){
	
	url:= context.Param("url")
	result, err := handler.urlService.GetShortenUrl(url)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.PageNotFound)
	}
	context.Redirect(http.StatusFound, result)
}

func(handler *V1UrlController)GetUrlStats(context *gin.Context){
	url:=context.Param("url")

	result, err := handler.urlService.GetUrlStats(url)
	if nil != err{
		handler.errorHelper.HTTPResponseError(context,err,http.StatusNotFound)
	}
	context.JSON(http.StatusOK,result)
}
