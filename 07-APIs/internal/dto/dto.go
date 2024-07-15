package dto

import "github.com/shopspring/decimal"

type CreateProductInput struct {
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
}
