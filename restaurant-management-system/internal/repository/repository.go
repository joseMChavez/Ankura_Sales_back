package repository

import (
    "database/sql"
    "errors"
)

type Restaurant struct {
    ID       int
    Name     string
    Location string
}

type Repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) CreateRestaurant(restaurant Restaurant) (int, error) {
    result, err := r.db.Exec("INSERT INTO restaurants (name, location) VALUES (?, ?)", restaurant.Name, restaurant.Location)
    if err != nil {
        return 0, err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
    return int(id), nil
}

func (r *Repository) GetRestaurant(id int) (Restaurant, error) {
    var restaurant Restaurant
    err := r.db.QueryRow("SELECT id, name, location FROM restaurants WHERE id = ?", id).Scan(&restaurant.ID, &restaurant.Name, &restaurant.Location)
    if err != nil {
        if err == sql.ErrNoRows {
            return Restaurant{}, errors.New("restaurant not found")
        }
        return Restaurant{}, err
    }
    return restaurant, nil
}

func (r *Repository) UpdateRestaurant(restaurant Restaurant) error {
    _, err := r.db.Exec("UPDATE restaurants SET name = ?, location = ? WHERE id = ?", restaurant.Name, restaurant.Location, restaurant.ID)
    return err
}

func (r *Repository) DeleteRestaurant(id int) error {
    _, err := r.db.Exec("DELETE FROM restaurants WHERE id = ?", id)
    return err
}