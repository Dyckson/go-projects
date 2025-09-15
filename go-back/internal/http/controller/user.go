package controller

import (
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
	var previewUser domain.User

	if err := c.ShouldBindJSON(&previewUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid request body",
		})
		return
	}

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

	if previewUser.Name != "" && previewUser.Name != currentUser.Name {
		currentUser.Name = previewUser.Name
	}
	if previewUser.Email != "" && previewUser.Email != currentUser.Email {
		currentUser.Email = previewUser.Email
	}

	user, err := uc.UserService.UpdateUser(currentUser)
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
	unwrapped := errors.Unwrap(err)

	if err != nil && !errors.Is(unwrapped, ErrNoRows) {
		log.Printf("controller=UserController func=CreateUser email=%s err=%v", input.Email, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "failed to create user",
		})
		return
	}

	if user.UUID == "" {
		user, err = uc.UserService.CreateUser(input)
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
			"message": "User created.",
			"data":    user,
		})
	}
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
