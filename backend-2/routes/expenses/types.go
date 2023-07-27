package expenses

import (
	"sapp/paperless-accounting/config"
)

type Expense struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type ExpenseRouter struct {
	conf *config.Config
}
