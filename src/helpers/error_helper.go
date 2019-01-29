package helpers

import (
	"../constants"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/rollbar/rollbar-go"
)

type ErrorHelper struct {
}

func ErrorHelperHandler() (ErrorHelper) {
	return ErrorHelper{}
}

func (handler *ErrorHelper) HTTPResponseError(context *gin.Context, e error, defaultErrorCode int) {

	switch e.Error() {
	case "record not found":
		defaultErrorCode = constants.ResourceNotFound
		break
	}

	errorConstant := constants.GetErrorConstant(defaultErrorCode)
	context.JSON(errorConstant.HttpCode, gin.H{
		"code":    defaultErrorCode,
		"message": errorConstant.Message,
	})

	if _, ok := e.(*mysql.MySQLError); !ok {
		rollbar.Error(e)
	}

	panic(e)

}
