package handler

import (
	"go-back/internal/http/controller"

	"github.com/gin-gonic/gin"
)

func HandleRequests(router *gin.Engine) {
	api := router.Group("/api")
	api.GET("/check", controller.HealthCheckStatus)

	userController := &controller.UserController{}

	user := api.Group("/user")
	user.GET("/list", userController.ListAllUsers)
	user.GET("/list/:userUUID", userController.ListUser)
	user.PUT("/edit/:userUUID", userController.UpdateUser)
	user.PUT("/delete/:userUUID", userController.DeleteUser)
	user.POST("/create", userController.CreateUser)
}
