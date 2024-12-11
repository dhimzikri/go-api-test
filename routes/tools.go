package routes

import (
	"go_api_cimol/controllers"
	"go_api_cimol/middleware"

	"github.com/gorilla/mux"
)

func ToolsRoutes(r *mux.Router) {
	router := r.PathPrefix("/tools").Subrouter()

	router.Use(middleware.Auth)
	router.Use(middleware.RateLimiter)

	router.HandleFunc("/Me", controllers.Me).Methods("GET")
	router.HandleFunc("/GetDataAssignFC", controllers.GetDataAssignFC).Methods("POST")
	router.HandleFunc("/FindDataAssignFC", controllers.FindDataAssignFC).Methods("POST")
	router.HandleFunc("/GetTblMappingResult", controllers.GetTblMappingResult).Methods("GET")
	router.HandleFunc("/GetTblQuestion", controllers.GetTblQuestion).Methods("POST")
	router.HandleFunc("/GetLovGroup", controllers.GetLovGroup).Methods("POST")
	router.HandleFunc("/GetCustomerCard", controllers.GetCustomerCard).Methods("POST")
	router.HandleFunc("/GetCollectionHistory", controllers.GetCollectionHistory).Methods("POST")
	router.HandleFunc("/GetDataAssignFcHistory", controllers.GetDataAssignFcHistory).Methods("POST")
	router.HandleFunc("/GetTblFieldCollHD", controllers.GetTblFieldCollHD).Methods("POST")
	router.HandleFunc("/FindDataAssignFcHistory", controllers.FindDataAssignFcHistory).Methods("POST")
	router.HandleFunc("/GetTblDataColl", controllers.GetTblDataColl).Methods("POST")
	router.HandleFunc("/GetTblResultData", controllers.GetTblResultData).Methods("POST")
}
