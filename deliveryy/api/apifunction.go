package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Exoplanet struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Distance    float64 `json:"distance"`
	Radius      float64 `json:"radius"`
	Mass        float64 `json:"mass,omitempty"`
	Type        string  `json:"type"`
}

var ExoplanetDatabase []Exoplanet

// AddExoplanetHandler adds a new exoplanet to the database
func AddExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	var newPlanet Exoplanet
	json.NewDecoder(r.Body).Decode(&newPlanet)
	ExoplanetDatabase = append(ExoplanetDatabase, newPlanet)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPlanet)
}

// ListExoplanetsHandler returns a list of all available exoplanets
func ListExoplanetsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ExoplanetDatabase)
}

// GetExoplanetByIDHandler retrieves information about a specific exoplanet by its ID
func GetExoplanetByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, planet := range ExoplanetDatabase {
		if planet.ID == params["id"] {
			json.NewEncoder(w).Encode(planet)
			return
		}
	}
	http.Error(w, "Exoplanet not found", http.StatusNotFound)
}

// UpdateExoplanetHandler updates the details of an existing exoplanet
func UpdateExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedPlanet Exoplanet
	json.NewDecoder(r.Body).Decode(&updatedPlanet)
	for i, planet := range ExoplanetDatabase {
		if planet.ID == params["id"] {
			ExoplanetDatabase[i] = updatedPlanet
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedPlanet)
			return
		}
	}
	http.Error(w, "Exoplanet not found", http.StatusNotFound)
}

// DeleteExoplanetHandler removes an exoplanet from the database
func DeleteExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, planet := range ExoplanetDatabase {
		if planet.ID == params["id"] {
			ExoplanetDatabase = append(ExoplanetDatabase[:i], ExoplanetDatabase[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Exoplanet with ID %s is deleted", params["id"])
			return
		}
	}
	http.Error(w, "Exoplanet not found", http.StatusNotFound)
}

// FuelEstimationHandler provides an overall fuel cost estimation for a trip to a particular exoplanet for a given crew capacity
func FuelEstimationHandler(w http.ResponseWriter, r *http.Request) {
	exoplanetID := r.URL.Query().Get("exoplanet_id")
	crewCapacity := r.URL.Query().Get("crew_capacity")

	// Your fuel estimation logic here

	// For now, let's just return a placeholder response
	response := fmt.Sprintf("Estimated fuel cost for a trip to exoplanet %s with crew capacity %s is 1000 units", exoplanetID, crewCapacity)
	fmt.Fprintf(w, response)
}
