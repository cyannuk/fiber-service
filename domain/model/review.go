package model

import (
	"time"
)

//go:generate reform

//reform:reviews
type Review struct {
	ID        int64     `reform:"id,pk"`
	CreatedAt time.Time `reform:"created_at"`
	Reviewer  string    `reform:"reviewer"`
	ProductID int64     `reform:"product_id"`
	Rating    int32     `reform:"rating"`
	Body      string    `reform:"body"`
}
