package userhandler

import (
	"log"
	"mmoz/crud/modules"
	"mmoz/crud/modules/user"
	"mmoz/crud/modules/user/userusecase"
	"mmoz/crud/validate"
	"net/http"
	"github.com/gin-gonic/gin"
)

type (
	UserHandlerService interface {
		GetUserAllUsers(c *gin.Context)
		CreateUser(c *gin.Context)
		GetUserByUsername(c *gin.Context)
	}
	userHandler struct {
		userUsecase userusecase.UserUsecaseService
	}
)

func NewUserHandler(userUsecase userusecase.UserUsecaseService) UserHandlerService {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (h *userHandler) GetUserAllUsers(c *gin.Context) {

	users, err := h.userUsecase.GetUserAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error getting users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   users,
	})
}

func (h *userHandler) CreateUser(c *gin.Context) {

	var req user.CreateUserReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("Error binding json: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error binding json",
		})
		return
	}

	err = validate.ValidateStruct(req)
	if err != nil {
		log.Printf("Error validating struct: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err = h.userUsecase.CreatePlayer(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User created",
	})
}

func (h *userHandler) GetUserByUsername(c *gin.Context) {

	username := c.Param("username")
	token := new(modules.Token)
	token.Username, _ = c.Get("user")
	token.Role, _ = c.Get("role")

	user, err := h.userUsecase.GetUserByUsername(username, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   user,
	})
}
