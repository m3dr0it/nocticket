package handler

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"noctiket/model/entity"
	"noctiket/model/request"
	"noctiket/repository"
	"noctiket/response"
	"noctiket/util"
)

func AddRolePermission(c *gin.Context) {
	userInterface, exist := c.Get("user")
	user, ok := userInterface.(entity.User)

	if !exist || !ok {
		response.UnauthorizedResponse(c)
		return
	}

	if user.Role != "admin" {
		response.UnauthorizedResponse(c)
		return
	}

	var permission entity.RolePermission
	err := c.ShouldBindJSON(&permission)
	if err != nil {
		response.ErrorInvalidRequest(c)
		return
	}

	err = repository.AddRolePermission(permission)
	if err != nil {
		response.MapResponseByError(c, err)
		return
	}

	response.SuccessResponse(c, nil)
}

func GetAllRolePermission(c *gin.Context) {
	userInterface, exist := c.Get("user")
	_, ok := userInterface.(entity.User)

	if !exist || !ok {
		response.UnauthorizedResponse(c)
		return
	}

	res, err := repository.GetRolePermission("")
	if err != nil {
		response.MapResponseByError(c, err)
		return
	}

	if len(res) < 1 {
		response.SuccessResponse(c, nil)
		return
	}

	response.SuccessResponse(c, res)
}

func DeleteRolePermission(c *gin.Context) {
	authorizationToken := c.Request.Header.Get("Authorization")
	claims, err := util.GetClaims(authorizationToken)
	if err != nil {
		response.UnauthorizedResponse(c)
		return
	}

	if claims.Role != "admin" {
		response.UnauthorizedResponse(c)
		return
	}

	var deletePermission request.DeleteRolePermission
	err = c.ShouldBindJSON(&deletePermission)
	if err != nil {
		response.ErrorInvalidRequest(c)
	}

	res, err := repository.GetRolePermissionById(deletePermission.Id)
	if err != nil {
		response.MapResponseByError(c, err)
		return
	}

	if len(res) < 1 {
		response.MapResponseByError(c, mongo.ErrNoDocuments)
		return
	}

	err = repository.DeleteRolePermission(deletePermission.Id)
	if err != nil {
		response.MapResponseByError(c, err)
		return
	}

	response.SuccessResponse(c, nil)
}
