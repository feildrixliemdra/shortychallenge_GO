package constants

import "net/http"

type ErrorConstant struct {
	HttpCode int
	Message  string
}

const (
	RequestParameterInvalid      = 1000
	ObjectNotInitializedProperly = 1001
	InternalServerError          = 1002
	ResourceNotFound             = 1003
	PageNotFound                 = 404
	StatusUnprocessableEntity    = 422
)

var errorConstantMapping = map[int]ErrorConstant{

	RequestParameterInvalid: {
		HttpCode: http.StatusBadRequest,
		Message:  "Invalid request parameter",
	},

	ObjectNotInitializedProperly: {
		HttpCode: http.StatusInternalServerError,
		Message:  "Object is not initialized properly",
	},

	InternalServerError: {
		HttpCode: http.StatusInternalServerError,
		Message:  "Something went wrong",
	},

	ResourceNotFound: {
		HttpCode: http.StatusInternalServerError,
		Message:  "Resource not found",
	},
	PageNotFound: {
		HttpCode: http.StatusNotFound,
		Message:  "The shortcode cannot be found in the system",
	},
	StatusUnprocessableEntity: {
		HttpCode: http.StatusUnprocessableEntity,
		Message:  "Unprocessable Entity",
	},
}

func GetErrorConstant(code int) ErrorConstant {
	return errorConstantMapping[code]
}
