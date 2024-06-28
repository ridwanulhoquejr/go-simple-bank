package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ridwanulhoquejr/go-simple-bank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func NewServer(store *db.Store) *Server {
	srv := &Server{store: store}
	rtr := gin.Default()

	// add routers to router
	rtr.POST("/accounts", srv.createAccount)

	srv.router = rtr

	return srv
}
