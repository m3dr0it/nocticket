package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"noctiket/model/entity"
	"noctiket/model/request"
	"noctiket/repository"
	"noctiket/response"
)

func GetUser(c *gin.Context) {
	var userRequest request.UserRequest
	err := c.ShouldBindQuery(&userRequest)

	if err != nil {
		response.ErrorInvalidRequest(c)
	}

	results, err := repository.FindUsers(userRequest)

	if err != nil {
		log.Println(err.Error())
		response.MapResponseByError(c, err)
		return
	}

	if len(results) < 1 {
		response.SuccessResponse(c, nil)
		return
	}

	var responseData []gin.H
	for _, result := range results {
		responseData = append(responseData, gin.H(result))
	}
	response.SuccessResponse(c, responseData)
}

func Register(c *gin.Context) {
	var user entity.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		response.ErrorInvalidRequest(c)
		return
	}

	if !isEmailValid(user.Email) {
		response.ErrorInvalidEmail(c)
		return
	}

	if err = repository.AddUser(user); err != nil {
		log.Println(err.Error())
		response.MapResponseByError(c, err)
		return
	}

	response.SuccessResponse(c, nil)
}
