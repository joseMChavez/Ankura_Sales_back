package handlers

import (
	"encoding/json"
	"net/http"
	"restaurant-management-system/internal/models"
	"restaurant-management-system/internal/repository"

	"github.com/gorilla/mux"
)

// Restaurant represents the structure of a restaurant
type RestaurantHandler struct {
	Repo repository.RestaurantRepository
}

func NewRestaurantHandler(service repository.RestaurantRepository) *RestaurantHandler {
	return &RestaurantHandler{Repo: service}
}

// GetRestaurants handles the request to get all restaurants
func (rh *RestaurantHandler) GetRestaurants(w http.ResponseWriter, r *http.Request) {

	restaurants, err := rh.Repo.GetAll()
	if err != nil {
		http.Error(w, "Error al obtener los restaurantes", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)
}

// GetRestaurants handles the request to get one restaurants
func (rh *RestaurantHandler) GetRestaurant(w http.ResponseWriter, r *http.Request) {

	idStr := mux.Vars(r)["id"]

	if idStr == "" {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	restaurants, err := rh.Repo.GetByID(idStr)
	if err != nil {
		http.Error(w, "Error al obtener los restaurantes", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)
}

// CreateRestaurant handles the request to create a new restaurant
func (repo *RestaurantHandler) CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	var restaurant models.Restaurant
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if restaurant.Name == "" {
		http.Error(w, "El nombre del restaurante es obligatorio", http.StatusBadRequest)
		return
	}

	if err := repo.Repo.Create(restaurant); err != nil {
		http.Error(w, "Error al crear el restaurante: "+restaurant.Name, http.StatusInternalServerError)
		return
	}
	// Here you would typically save the restaurant to the database

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(restaurant)
}

// CreateRestaurant handles the request to create a new restaurant
func (repo *RestaurantHandler) UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	var restaurant models.Restaurant
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if restaurant.Name == "" {
		http.Error(w, "El nombre del restaurante es obligatorio", http.StatusBadRequest)
		return
	}

	if err := repo.Repo.Update(restaurant); err != nil {
		http.Error(w, "Error al crear el restaurante: "+restaurant.Name, http.StatusInternalServerError)
		return
	}
	// Here you would typically save the restaurant to the database

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(restaurant)
}
