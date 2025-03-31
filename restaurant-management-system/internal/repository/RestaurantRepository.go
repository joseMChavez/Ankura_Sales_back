package repository

import (
	"restaurant-management-system/internal/models"

	"gorm.io/gorm"
)

// RestaurantRepository define los métodos que deben implementarse para interactuar con la base de datos
type RestaurantRepository interface {
	Create(restaurant models.Restaurant) error
	GetByID(id string) (models.Restaurant, error)
	Update(restaurant models.Restaurant) error
	Delete(id string) error
}

// RestaurantRepositoryImpl es la implementación de la interfaz RestaurantRepository
type RestaurantRepositoryImpl struct {
	DB *gorm.DB
}

// NewRestaurantRepository crea una nueva instancia de RestaurantRepositoryImpl
func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {
	return &RestaurantRepositoryImpl{DB: db}
}

// Create guarda un nuevo restaurante en la base de datos
func (r *RestaurantRepositoryImpl) Create(restaurant models.Restaurant) error {
	return r.DB.Create(&restaurant).Error
}

// GetByID obtiene un restaurante por su ID
func (r *RestaurantRepositoryImpl) GetByID(id string) (models.Restaurant, error) {
	var restaurant models.Restaurant
	err := r.DB.First(&restaurant, "id = ?", id).Error
	return restaurant, err
}

// Update actualiza un restaurante existente en la base de datos
func (r *RestaurantRepositoryImpl) Update(restaurant models.Restaurant) error {
	return r.DB.Save(&restaurant).Error
}

// Delete elimina un restaurante por su ID
func (r *RestaurantRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&models.Restaurant{}, "id = ?", id).Error
}
