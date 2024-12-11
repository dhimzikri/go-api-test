package routes

import (
	"go_api_cimol/controllers"
	"go_api_cimol/middleware"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()
	router.Use(middleware.RateLimiter)

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")
}
