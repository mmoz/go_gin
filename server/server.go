package server

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type (
	Server struct {
		Router *gin.Engine
		db     *sql.DB
	}
)

func StartServer(db *sql.DB) *Server {
	return &Server{
		Router: gin.Default(),
		db:     db,
	}
}
