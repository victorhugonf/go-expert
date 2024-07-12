package entity

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
	"github.com/victorhugonf/go-expert/07-APIs/pkg/entity"
)

var (
	ErrRequiredId    = errors.New("id is required")
	ErrInvalidId     = errors.New("id is invalid")
	ErrRequiredName  = errors.New("name is required")
	ErrRequiredPrice = errors.New("price is required")
	ErrInvalidPrice  = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID       `json:"id"`
	Name      string          `json:"name"`
	Price     decimal.Decimal `json:"price"`
	CreatedAt time.Time       `json:"createdAt"`
}

func NewProduct(name string, price decimal.Decimal) (*Product, error) {
	product := &Product{
		ID:        entity.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrRequiredId
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidId
	}
	if p.Name == "" {
		return ErrRequiredName
	}
	if p.Price.Compare(decimal.NewFromInt(0)) == 0 {
		return ErrRequiredPrice
	}
	if p.Price.Compare(decimal.NewFromInt(0)) == -1 {
		return ErrInvalidPrice
	}
	return nil
}
