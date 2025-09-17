package controller

import (
	"database/sql"
	"errors"
	"go-back/internal/domain"
	"go-back/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrNoRows = errors.New("sql: no rows in result set")

type UserController struct {
	UserService service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{UserService: s}
}

func (uc *UserController) ListAllUsers(c *gin.Context) {
	users, err := uc.UserService.ListAllUsers()
	if err != nil {
		log.Printf("controller=UserController func=ListUser err=%v", err)

		status := http.StatusInternalServerError
		message := "internal error"

		if errors.Is(err, ErrNoRows) {
			status = http.StatusNotFound
			message = "no user found for this userUUID"
		}

		c.AbortWithStatusJSON(status, gin.H{
			"success": false,
			"message": message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}

func (uc *UserController) ListUser(c *gin.Context) {
	userUUID := c.Param("userUUID")

	user, err := uc.UserService.ListUserByUUID(userUUID)
	if err != nil {
		log.Printf("controller=UserController func=ListUser userUUID=%s err=%v", userUUID, err)

		status := http.StatusInternalServerError
		message := "internal error"

		if errors.Is(err, ErrNoRows) {
			status = http.StatusNotFound
			message = "no user found for this userUUID"
		}

		c.AbortWithStatusJSON(status, gin.H{
			"success": false,
			"message": message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	userUUID := c.Param("userUUID")
	if userUUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "userUUID is required",
		})
		return
	}

	var previewUser domain.User
	if err := c.ShouldBindJSON(&previewUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
		})
		return
	}
	previewUser.UUID = userUUID

	currentUser, err := uc.UserService.ListUserByUUID(previewUser.UUID)
	if err != nil {
		log.Printf("controller=UserController func=UpdateUser userUUID=%s err=%v", previewUser.UUID, err)
		status := http.StatusInternalServerError
		message := "internal error"
		if errors.Is(err, ErrNoRows) {
			status = http.StatusNotFound
			message = "no user found for this userUUID"
		}
		c.AbortWithStatusJSON(status, gin.H{
			"success": false,
			"message": message,
		})
		return
	}

	changed := false

	if previewUser.Name != "" && previewUser.Name != currentUser.Name {
		currentUser.Name = previewUser.Name
		changed = true
	}

	if previewUser.Email != "" && previewUser.Email != currentUser.Email {
		currentUser.Email = previewUser.Email
		changed = true
	}

	if !changed {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "No changes detected.",
			"data":    currentUser,
		})
		return
	}

	updatedUser, err := uc.UserService.UpdateUser(currentUser)
	if err != nil {
		log.Printf("controller=UserController func=UpdateUser userUUID=%s err=%v", previewUser.UUID, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User updated.",
		"data":    updatedUser,
	})
}

func (uc *UserController) ManageActivateUser(c *gin.Context) {
	userUUID := c.Param("userUUID")
	if userUUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "userUUID is required",
		})
		return
	}

	user, err := uc.UserService.ManageActivateUser(userUUID)
	if err != nil {
		log.Printf("controller=UserController func=UpdateUser userUUID=%s err=%v", userUUID, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User updated.",
		"data":    user,
	})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var input domain.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
		})
		return
	}

	user, err := uc.UserService.ListUserByEmail(input.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Printf("controller=UserController func=CreateUser email=%s err=%v", input.Email, err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "failed to check user",
			})
			return
		}
		user = domain.User{}
	}

	if user.UUID != "" {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "user already exists",
			"data":    user,
		})
		return
	}

	newUser, err := uc.UserService.CreateUser(input)
	if err != nil {
		log.Printf("controller=UserController func=CreateUser email=%s err=%v", input.Email, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "user created",
		"data":    newUser,
	})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userUUID := c.Param("userUUID")

	err := uc.UserService.DeleteUser(userUUID)
	if err != nil {
		log.Printf("controller=UserController func=DeleteUser userUUID=%s err=%v", userUUID, err)
		status := http.StatusInternalServerError
		message := "internal error"

		c.AbortWithStatusJSON(status, gin.H{
			"success": false,
			"message": message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User deleted successfully",
	})
}
