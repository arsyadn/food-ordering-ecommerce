package models

import "time"

type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `json:"user_id"`
	Total     float64     `json:"total"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Items     []OrderItem `json:"items"` // relasi ke OrderItem
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	MenuID    uint    `json:"menu_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`     // harga satuan
	Subtotal  float64 `json:"subtotal"`  // qty * price
}
