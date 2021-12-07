package router

import (
	"net/http"
	"solaris/middleware"

	"github.com/gorilla/mux"
)

// Router is exported to be used in main
func Router() *mux.Router {

	router := mux.NewRouter()

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	// Web App
	router.PathPrefix("/assets").Handler(staticFileHandler).Methods("GET")
	//router.handleFunc("/", main.).Methods("GET")

	// APIs
	// Electric Meter
	router.HandleFunc("/api/electricmeter/{id}", middleware.GetElectricMeter).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/electricmeter", middleware.GetAllElectricMeters).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newelectricmeter", middleware.CreateElectricMeter).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/electricmeter/{id}", middleware.UpdateElectricMeter).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteelectricmeter/{id}", middleware.DeleteElectricMeter).Methods("DELETE", "OPTIONS")

	// Inverter
	router.HandleFunc("/api/inverter/{id}", middleware.GetInverter).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/inverter", middleware.GetAllInverters).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newinverter", middleware.CreateInverter).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/inverter/{id}", middleware.UpdateInverter).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteinverter/{id}", middleware.DeleteInverter).Methods("DELETE", "OPTIONS")

	return router
}
