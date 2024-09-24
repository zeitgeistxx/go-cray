package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeitgeistxx/go-api/services/cart"
	"github.com/zeitgeistxx/go-api/services/order"
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
	cartSubrouter := router.PathPrefix("/api/v1/cart").Subrouter()

	userStore := user.NewStore(s.db)
	productStore := product.NewStore(s.db)
	orderStore := order.NewStore(s.db)

	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(userSubrouter)

	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(productSubrouter)

	cartHandler := cart.NewHandler(productStore, orderStore, userStore)
	cartHandler.RegisterRoutes(cartSubrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
