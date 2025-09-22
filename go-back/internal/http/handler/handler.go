package handler

import (
	"go-back/internal/http/controller"
	"go-back/internal/service"
	"go-back/internal/storage/repository"

	"github.com/gin-gonic/gin"
)

func HandleRequests(router *gin.Engine) {
	api := router.Group("/api")
	api.GET("/check", controller.HealthCheckStatus)

	repo := &repository.UserRepository{}
	userService := service.NewUserService(repo)
	userController := &controller.UserController{UserService: userService}

	user := api.Group("/user")
	user.GET("/list", userController.ListAllUsers)
	user.GET("/list/:userUUID", userController.ListUser)

	user.POST("/create", userController.CreateUser)

	user.PUT("/edit/:userUUID", userController.UpdateUser)
	user.PUT("/manage/:userUUID", userController.ManageActivateUser)

	user.DELETE("/delete/:userUUID", userController.DeleteUser)
}
