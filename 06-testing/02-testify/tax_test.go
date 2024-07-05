package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxWhenAmountEquals1000(t *testing.T) {
	actual, err := CalculateTax(1000.0)

	assert.Nil(t, err)
	assert.Equal(t, 10.0, actual)
}

func TestCalculateTaxWhenAmountIsNegative(t *testing.T) {
	actual, err := CalculateTax(-1)

	assert.EqualError(t, err, "amount must be grater than 0")
	assert.Equal(t, 0.0, actual)
}
