package server

import (
	"mmoz/crud/middleware"
	"mmoz/crud/modules/user/userhandler"
	"mmoz/crud/modules/user/userrepository"
	"mmoz/crud/modules/user/userusecase"
)

// func StartUserServer(r *gin.Engine, db *sql.DB) {
// 	userRepo := userrepository.NewUserRepository(db)
// 	userUsecase := userusecase.NewUserUsecase(userRepo)
// 	userhandler := userhandler.NewUserHandler(userUsecase)
// 	api := r.Group("/api")
// 	authen := r.Group("/api")
// 	authen.Use(middleware.AuthMiddleware())
// 	authen.GET("/users", userhandler.GetUserAllUsers)
// 	api.POST("/users", userhandler.CreateUser)
// 	authen.GET("/users/:username", userhandler.GetUserByUsername)

// }

func (s *Server) StartUserServer() {
	userRepo := userrepository.NewUserRepository(s.db)
	userUsecase := userusecase.NewUserUsecase(userRepo)
	userhandler := userhandler.NewUserHandler(userUsecase)
	api := s.Router.Group("/api")
	authen := s.Router.Group("/api")
	admin := s.Router.Group("/api")
	authen.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminMiddleware())
	admin.GET("/users", userhandler.GetUserAllUsers)
	api.POST("/users", userhandler.CreateUser)
	authen.GET("/users/:username", userhandler.GetUserByUsername)
}
