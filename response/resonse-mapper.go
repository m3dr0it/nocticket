package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
)

const (
	SuccessMessage             = "Request processed successfully"
	DuplicateEntryError        = "Duplicate entry detected"
	InternalServerError        = "Something went wrong, please try again later"
	InvalidRequestError        = "Invalid request"
	InvalidEmail               = "Invalid Email"
	DataNotFoundMessage        = "Data Not Found"
	WrongPasswordMessage       = "Wrong Password"
	UnauthorizedResponsMessage = "Unauthorized"
)

func SuccessResponse(c *gin.Context, data interface{}) {
	if data == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": SuccessMessage,
			"data":    []interface{}{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": SuccessMessage,
			"data":    data,
		})
	}

}

func errorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  "error",
		"message": message,
	})
}

func ErrorInvalidEmail(c *gin.Context) {
	errorResponse(c, 400, InvalidEmail)
}

func ErrorInvalidRequest(c *gin.Context) {
	errorResponse(c, 400, InvalidRequestError)
}

func WrongPasswordResponse(c *gin.Context) {
	errorResponse(c, 400, WrongPasswordMessage)
}

func UnauthorizedResponse(c *gin.Context) {
	errorResponse(c, 401, UnauthorizedResponsMessage)
}

func MapResponseByError(c *gin.Context, err error) {
	switch true {
	case strings.Contains(err.Error(), "duplicate"):
		errorResponse(c, 400, DuplicateEntryError)
		return
	case errors.Is(err, mongo.ErrNoDocuments):
		errorResponse(c, 404, DataNotFoundMessage)
		return
	default:
		errorResponse(c, 500, InternalServerError)
		return
	}
}
