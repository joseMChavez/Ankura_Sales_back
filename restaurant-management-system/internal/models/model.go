package models

type Restaurant struct {
    ID       string  `json:"id"`
    Name     string  `json:"name"`
    Location string  `json:"location"`
    Menu     []MenuItem `json:"menu"`
}

type MenuItem struct {
    ID       string  `json:"id"`
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Category string  `json:"category"`
}