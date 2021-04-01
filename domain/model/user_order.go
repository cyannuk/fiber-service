package model

//go:generate reform

//reform:user_orders
type UserOrder struct {
	Name      string  `reform:"name"`
	City      string  `reform:"city"`
	State     string  `reform:"state"`
	ProductID int64   `reform:"product_id"`
	Quantity  int32   `reform:"quantity"`
	Total     float64 `reform:"total"`
}
