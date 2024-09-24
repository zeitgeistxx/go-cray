package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeitgeistxx/go-api/services/product"
	"github.com/zeitgeistxx/go-api/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	
	userSubrouter := router.PathPrefix("/api/v1/users").Subrouter()
	productSubrouter := router.PathPrefix("/api/v1/products").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(userSubrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(productSubrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
