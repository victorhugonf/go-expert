package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TaxRespositoryMock struct {
	mock.Mock
}

func (m *TaxRespositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}

func TestCalculateAndSaveWhenAmountEquals1000(t *testing.T) {
	repository := &TaxRespositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)

	err := CalculateTaxAndSave(1000.0, repository)

	assert.Nil(t, err)
	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 1)
}

func TestCalculateAndSaveWhenAmountIsNegative(t *testing.T) {
	repository := &TaxRespositoryMock{}
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax")).Once()
	//repository.On("SaveTax", mock.Anything).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(-4.1, repository)

	assert.EqualError(t, err, "error saving tax")
	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 1)
}

func TestCalculateTaxWhenAmountEquals1000(t *testing.T) {
	actual := calculateTax(1000.0)

	assert.Equal(t, 10.0, actual)
}

func TestCalculateTaxWhenAmountIsNegative(t *testing.T) {
	actual := calculateTax(-1)

	assert.Equal(t, 0.0, actual)
}
