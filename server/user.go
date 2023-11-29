package server

import (
	"database/sql"
	"mmoz/crud/middleware"
	"mmoz/crud/modules/user/userhandler"
	"mmoz/crud/modules/user/userrepository"
	"mmoz/crud/modules/user/userusecase"

	"github.com/gin-gonic/gin"
)

func StartUserServer(r *gin.Engine, db *sql.DB) {
	userRepo := userrepository.NewUserRepository(db)
	userUsecase := userusecase.NewUserUsecase(userRepo)
	userhandler := userhandler.NewUserHandler(userUsecase)
	api := r.Group("/api")
	authen := r.Group("/api")
	authen.Use(middleware.AuthMiddleware())
	authen.GET("/users", userhandler.GetUserAllUsers)
	api.POST("/users", userhandler.CreateUser)
	authen.GET("/users/:username", userhandler.GetUserByUsername)

}
