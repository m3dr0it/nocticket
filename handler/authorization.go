package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"noctiket/model/entity"
	"noctiket/model/request"
	"noctiket/repository"
	"noctiket/response"
	"noctiket/util"
)

func Authorization(c *gin.Context) {
	authorizationToken := c.Request.Header.Get("Authorization")
	user, err := getLoggedUser(authorizationToken)

	if err != nil {
		response.UnauthorizedResponse(c)
		c.Abort()
		return
	}

	c.Set("user", user)
	userInterface, exist := c.Get("user")
	user, ok := userInterface.(entity.User)

	if !exist || !ok {
		response.UnauthorizedResponse(c)
		c.Abort()
		return
	}

	if user.Role == "admin" {
		c.Next()
		return
	}

	res, err := repository.GetRolePermission(user.Role)
	if res == nil || err != nil {
		response.UnauthorizedResponse(c)
		c.Abort()
		return
	}

	if len(res) < 1 {
		response.UnauthorizedResponse(c)
		c.Abort()
		return
	}

	var rolePermission entity.RolePermission
	rpBytes, err := bson.Marshal(res[0])
	_ = bson.Unmarshal(rpBytes, &rolePermission)

	isAuthorized := false

	for _, v := range rolePermission.Apis {
		log.Println(c.FullPath())
		if v.Path == c.FullPath() && v.Method == c.Request.Method {
			isAuthorized = true
		}
	}

	if !isAuthorized {
		response.UnauthorizedResponse(c)
		c.Abort()
		return
	}

	c.Next()
}

func Login(c *gin.Context) {
	var login request.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		response.ErrorInvalidRequest(c)
		return
	}

	if !isEmailValid(login.Email) {
		response.ErrorInvalidEmail(c)
		return
	}

	user, err := repository.FindUserByEmail(login.Email)

	if err != nil {
		log.Println(err.Error())
		response.MapResponseByError(c, err)
		return
	}

	log.Println(user.Id)

	if user.Password != util.MD5Hash(login.Password) {
		log.Println("Wrong password")
		response.WrongPasswordResponse(c)
		return
	}

	response.SuccessResponse(c, response.TokenResponse{
		Token: util.GenerateJWT(user),
	})
}

func getLoggedUser(authorization string) (entity.User, error) {
	claims, err := util.GetClaims(authorization)

	if err != nil {
		return entity.User{}, errors.New("token Unavailable")
	}

	if claims.Email == "" {
		return entity.User{}, err
	}

	user, err := repository.FindUserByEmail(claims.Email)
	if err != nil {
		log.Println(err.Error())
		return entity.User{}, err
	}

	return user, nil
}
