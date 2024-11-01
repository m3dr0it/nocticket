package main

import (
	"github.com/gin-gonic/gin"
	"noctiket/handler"
	"noctiket/repository"
)

func main() {
	repository.InitDbConnection("nocticket")
	router := gin.Default()

	router.POST("/open/api/v1/register", handler.Register)
	router.POST("/open/api/v1/login", handler.Login)
	api := router.Group("/api", handler.Authorization)
	apiV1 := api.Group("/v1")
	apiV1.GET("/user", handler.GetUser)
	apiV1.POST("/ticket", handler.CreateTicket)
	apiV1.GET("/ticket", handler.GetTickets)
	apiV1.POST("/ticket/assign", handler.AssignTicket)
	apiV1.GET("/engineer/ticket", handler.GetTickets)
	apiV1.POST("/admin/permission", handler.AddRolePermission)
	apiV1.POST("/ticket/progress", handler.UpdateProgress)
	apiV1.GET("/admin/permission", handler.GetAllRolePermission)
	apiV1.DELETE("/admin/permission", handler.DeleteRolePermission)

	err := router.Run(":8000")
	if err != nil {
		return
	}
}
