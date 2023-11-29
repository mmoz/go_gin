package authhandler

import (
	"log"
	"mmoz/crud/modules/auth"
	"mmoz/crud/modules/auth/authusecase"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.authUsecase.CheckLogin(user)
	if err != nil {
		log.Printf("Error checking login: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)

}
