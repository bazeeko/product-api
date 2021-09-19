package server

import (
	"log"
	"net/http"
	"simple-api/config"

	"github.com/gorilla/mux"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(conf *config.Config, router *mux.Router) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    conf.Host + conf.Port,
			Handler: router,
		},
	}
}

func (server *Server) Run() error {
	log.Println("API is running at", server.httpServer.Addr)
	return server.httpServer.ListenAndServe()
}
