package entity

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Beer", decimal.NewFromFloat(5.7))

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Beer", p.Name)
	assert.Equal(t, decimal.NewFromFloat(5.7), p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", decimal.NewFromFloat(5.8))

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrRequiredName, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Beer", decimal.NewFromInt(0))

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrRequiredPrice, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Beer", decimal.NewFromFloat(-10.9))

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Beer", decimal.NewFromFloat(3.1))

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
