package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sbercloud_test/pkg/handler"
)

type Server struct {
	router  *mux.Router
	handler *handler.Handler
}

func NewServer(h *handler.Handler) *Server {
	return &Server{router: mux.NewRouter(), handler: h}
}

func (server *Server) Run() {
	server.handler.RegisterHandlers(server.router)
	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	log.Fatal(http.ListenAndServe(":3228", handlers.CORS(originsOk, headersOk, methodsOk)(server.router)))
}
