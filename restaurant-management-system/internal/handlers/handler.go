package handlers

import (
    "net/http"
    "encoding/json"
)

// Restaurant represents the structure of a restaurant
type Restaurant struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Address string `json:"address"`
}

// GetRestaurants handles the request to get all restaurants
func GetRestaurants(w http.ResponseWriter, r *http.Request) {
    // Sample data
    restaurants := []Restaurant{
        {ID: "1", Name: "Restaurant A", Address: "123 Main St"},
        {ID: "2", Name: "Restaurant B", Address: "456 Elm St"},
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(restaurants)
}

// CreateRestaurant handles the request to create a new restaurant
func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
    var restaurant Restaurant
    if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Here you would typically save the restaurant to the database

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(restaurant)
}