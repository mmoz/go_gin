package authhandler

import (
	"log"
	"mmoz/crud/modules/auth"
	"mmoz/crud/modules/auth/authusecase"
	"mmoz/crud/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	AuthHandlerService interface {
		Login(c *gin.Context)
	}
	authHandler struct {
		authUsecase authusecase.AuthUsecaseService
	}
)

func NewAuthHandler(authUsecase authusecase.AuthUsecaseService) AuthHandlerService {
	return &authHandler{
		authUsecase: authUsecase,
	}
}

func (h *authHandler) Login(c *gin.Context) {

	user := new(auth.CredentialReq)
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding json: %v", err)
		response.ErrResponse(c, http.StatusBadRequest, "Error binding json")
		return
	}

	result, err := h.authUsecase.CheckLogin(user)
	if err != nil {
		log.Printf("Error checking login: %v", err)
		response.ErrResponse(c, http.StatusInternalServerError, err.Error())

	}

	response.SuccessResponse(c, http.StatusOK, result)
}
