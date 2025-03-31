package services

import (
	"errors"
	"restaurant-management-system/internal/models"
	"restaurant-management-system/internal/repository"
)

type RestaurantService struct {
	repo repository.RestaurantRepository
}

func NewRestaurantService(repo repository.RestaurantRepository) *RestaurantService {
	return &RestaurantService{repo: repo}
}

func (s *RestaurantService) CreateRestaurant(restaurant models.Restaurant) error {
	if restaurant.Name == "" {
		return errors.New("restaurant name cannot be empty")
	}
	return s.repo.Create(restaurant)
}

func (s *RestaurantService) GetRestaurant(id string) (models.Restaurant, error) {
	return s.repo.GetByID(id)
}

func (s *RestaurantService) UpdateRestaurant(restaurant models.Restaurant) error {
	if restaurant.ID == "" {
		return errors.New("restaurant ID cannot be empty")
	}
	return s.repo.Update(restaurant)
}

func (s *RestaurantService) DeleteRestaurant(id string) error {
	return s.repo.Delete(id)
}
