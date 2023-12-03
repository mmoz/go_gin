package userhandler

import (
	"log"
	"mmoz/crud/modules"
	"mmoz/crud/modules/user"
	"mmoz/crud/modules/user/userusecase"
	"mmoz/crud/pkg/request"
	"mmoz/crud/pkg/response"
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
		log.Printf("Error getting all users: %v", err)
		response.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, users)

}

func (h *userHandler) CreateUser(c *gin.Context) {

	req := new(user.CreateUserReq)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("Error binding json: %v", err)
		response.ErrResponse(c, http.StatusBadRequest, "Error binding json")
		return
	}

	err = request.ValidateStruct(req)
	if err != nil {
		log.Printf("Error validating struct: %v", err)
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.userUsecase.CreatePlayer(req)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		response.ErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "User created successfully")
}

func (h *userHandler) GetUserByUsername(c *gin.Context) {

	token := new(modules.Token)
	username := c.Param("username")
	token.Role, _ = c.Get("role")

	user, err := h.userUsecase.GetUserByUsername(username, token)
	if err != nil {
		log.Printf("Error getting user by username: %v", err)
		response.ErrResponse(c, http.StatusForbidden, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, user)
}
