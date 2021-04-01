package model

import (
	"time"
)

//go:generate reform

//reform:orders
type Order struct {
	ID        int64     `reform:"id,pk" json:"id,omit-dec"`
	CreatedAt time.Time `reform:"created_at" json:"created_at"`
	UserID    int64     `reform:"user_id" json:"user_id,omit-dec"`
	ProductID int64     `reform:"product_id" json:"product_id"`
	Discount  *float64  `reform:"discount" json:"discount"`
	Quantity  int32     `reform:"quantity" json:"quantity"`
	Subtotal  float64   `reform:"subtotal" json:"subtotal"`
	Tax       float64   `reform:"tax" json:"tax"`
	Total     float64   `reform:"total" json:"total"`
}
