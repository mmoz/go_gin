package server

import (
	"database/sql"
	"mmoz/crud/modules/auth/authhandler"
	"mmoz/crud/modules/auth/authrepository"
	"mmoz/crud/modules/auth/authusecase"

	"github.com/gin-gonic/gin"
)

func StartAuthServer(r *gin.Engine, db *sql.DB) {
	authRepo := authrepository.NewAuthRepository(db)
	authUsecase := authusecase.NewAuthUsecase(authRepo)
	authHandler := authhandler.NewAuthHandler(authUsecase)
	api := r.Group("/api")
	api.POST("/login", authHandler.Login)
}
