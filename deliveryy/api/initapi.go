package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func apiIntialize() {
	router := mux.NewRouter()
	router.HandleFunc("/exoplanets", AddExoplanetHandler).Methods("POST")
	router.HandleFunc("/exoplanets", ListExoplanetsHandler).Methods("GET")
	router.HandleFunc("/exoplanets/{id}", GetExoplanetByIDHandler).Methods("GET")
	router.HandleFunc("/exoplanets/{id}", UpdateExoplanetHandler).Methods("PUT")
	router.HandleFunc("/exoplanets/{id}", DeleteExoplanetHandler).Methods("DELETE")
	router.HandleFunc("/fuel-estimation", FuelEstimationHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", router))
}
