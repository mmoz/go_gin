package server

import (
	"mmoz/crud/modules/auth/authhandler"
	"mmoz/crud/modules/auth/authrepository"
	"mmoz/crud/modules/auth/authusecase"
)

// func StartAuthServer(r *gin.Engine, db *sql.DB) {
// 	authRepo := authrepository.NewAuthRepository(db)
// 	authUsecase := authusecase.NewAuthUsecase(authRepo)
// 	authHandler := authhandler.NewAuthHandler(authUsecase)
// 	api := r.Group("/api")
// 	api.POST("/login", authHandler.Login)
// }

func (s *Server) StartAuthServer() {
	authRepo := authrepository.NewAuthRepository(s.db)
	authUsecase := authusecase.NewAuthUsecase(authRepo)
	authHandler := authhandler.NewAuthHandler(authUsecase)
	api := s.Router.Group("/api")
	api.POST("/login", authHandler.Login)
}
