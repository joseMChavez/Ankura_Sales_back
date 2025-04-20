package models

import "time"

// Tabla: restaurants// Tabla: kitchen

// Tabla: restaurants
type Restaurant struct {
	ID       string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string    `json:"name" gorm:"type:varchar(255);not null"`
	Location string    `json:"location" gorm:"type:text;not null"`
	Kitchens []Kitchen `json:"kitchens" gorm:"foreignKey:RestaurantID"`
}

// Tabla: kitchen
type Kitchen struct {
	ID           int    `json:"id" gorm:"primaryKey;autoIncrement"`
	RestaurantID string `json:"restaurant_id" gorm:"type:uuid;not null"`
	Name         string `json:"name" gorm:"type:varchar(255);not null"`
	Description  string `json:"description" gorm:"type:text"`
	Menues       []Menu `json:"menues" gorm:"foreignKey:KitchenID"`
}

// Tabla: menues
type Menu struct {
	ID        string     `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	KitchenID int        `json:"kitchen_id" gorm:"not null"`
	Name      string     `json:"name" gorm:"type:varchar(255);not null"`
	Items     []MenuItem `json:"items" gorm:"foreignKey:MenuID"`
}

// Tabla: menu_items
type MenuItem struct {
	ID        string `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ImageURL  string `json:"image_url" gorm:"type:text;not null"`
	ProductID int    `json:"product_id" gorm:"not null"`
	MenuID    string `json:"menu_id" gorm:"type:uuid;not null"`
}

// Tabla: orders.tables
type Table struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Number string `json:"number" gorm:"type:varchar(50);unique;not null"`
	Name   string `json:"name" gorm:"type:varchar(100)"`
	State  string `json:"state" gorm:"type:varchar(50);default:'disponible'"`
}

// Tabla: factory.categoies_product
type CategoryProduct struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(100);unique;not null"`
}

// Tabla: factory.product
type Product struct {
	ID         int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string  `json:"name" gorm:"type:varchar(100);not null"`
	CategoryID int     `json:"category_id" gorm:"not null"`
	Price      float64 `json:"price" gorm:"type:numeric(10,2);not null"`
	Time       int     `json:"time" gorm:"not null"`
	LaborCost  float64 `json:"labor_cost" gorm:"type:numeric(10,2);not null"`
	IVA        float64 `json:"iva" gorm:"type:numeric(5,2);default:0.00"`
}

// Tabla: Inventories.Ingredients
type Ingredient struct {
	ID          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string  `json:"name" gorm:"type:varchar(100);unique;not null"`
	UnitCost    float64 `json:"unit_cost" gorm:"type:numeric(10,2);not null"`
	Unit        string  `json:"unit" gorm:"type:varchar(50);not null"`
	StockActual float64 `json:"stock_actual" gorm:"type:numeric(10,2);not null"`
	StockMinimo float64 `json:"stock_minimo" gorm:"type:numeric(10,2);not null"`
}

// Tabla: factory.recipes
type Recipe struct {
	ID           int     `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID    int     `json:"product_id" gorm:"not null"`
	IngredientID int     `json:"ingredient_id" gorm:"not null"`
	Quantity     float64 `json:"quantity" gorm:"type:numeric(10,2);not null"`
}

// Tabla: sales.billing
type Billing struct {
	ID      int     `json:"id" gorm:"primaryKey;autoIncrement"`
	TableID *int    `json:"table_id" gorm:"default:null"`
	State   string  `json:"state" gorm:"type:varchar(50);default:'pendiente'"`
	Total   float64 `json:"total" gorm:"type:numeric(10,2)"`
}

// Tabla: orders.billing_orders
type BillingOrder struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	BillingID int       `json:"billing_id" gorm:"not null"`
	State     string    `json:"state" gorm:"type:varchar(50);default:'pendiente'"`
	Total     float64   `json:"total" gorm:"type:numeric(10,2)"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
	Type      string    `json:"type" gorm:"type:varchar(50);not null;check:type IN ('dentro', 'fuera')"`
}

// Tabla: orders.billing_orders_det
type BillingOrderDetail struct {
	ID        int     `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID   int     `json:"order_id" gorm:"not null"`
	ProductID int     `json:"product_id" gorm:"not null"`
	Count     int     `json:"count" gorm:"not null"`
	Subtotal  float64 `json:"subtotal" gorm:"type:numeric(10,2);not null"`
}

// Tabla: sales.payment_method
type PaymentMethod struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(100);unique;not null"`
}

// Tabla: sales.invoices
type Invoice struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	BillID    int       `json:"bill_id" gorm:"not null"`
	MethodID  int       `json:"method_id" gorm:"not null"`
	Total     float64   `json:"total" gorm:"type:numeric(10,2);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
}

// Tabla: sales.payments
type Payment struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Total     float64   `json:"total" gorm:"type:numeric(10,2);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
}

// Tabla: sales.payment_invoices
type PaymentInvoice struct {
	ID        int `json:"id" gorm:"primaryKey;autoIncrement"`
	PaymentID int `json:"payment_id" gorm:"not null"`
	InvoiceID int `json:"invoice_id" gorm:"not null"`
}
