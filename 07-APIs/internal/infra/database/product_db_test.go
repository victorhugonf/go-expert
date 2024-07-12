package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/victorhugonf/go-expert/07-APIs/internal/entity"
)

func TestCreateProduct(t *testing.T) {
	db := DBOpen(t)
	db.AutoMigrate(&entity.Product{})
	productDB := NewProduct(db)
	product, _ := entity.NewProduct("Bike", decimal.NewFromFloat(123.4))

	err := productDB.Create(product)

	var productFound entity.Product
	assert.Nil(t, err)
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.NotNil(t, productFound.CreatedAt)
}

func TestFindAllProducts(t *testing.T) {
	db := DBOpen(t)
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), decimal.NewFromFloat(rand.Float64()*100))
		assert.NoError(t, err)
		db.Create(product)
	}
	productDB := NewProduct(db)

	products, err := productDB.FindAll(1, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindProductById(t *testing.T) {
	db := DBOpen(t)
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product", decimal.NewFromFloat(64))
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)

	productFound, err := productDB.FindByID(product.ID.String())

	assert.NoError(t, err)
	assert.Equal(t, "Product", productFound.Name)
}

func TestUpdateProduct(t *testing.T) {
	db := DBOpen(t)
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product", decimal.NewFromFloat(64))
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	product.Name = "New Product"

	err = productDB.Update(product)

	assert.NoError(t, err)
	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "New Product", productFound.Name)
}

func TestDeleteProduct(t *testing.T) {
	db := DBOpen(t)
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product", decimal.NewFromFloat(64))
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)

	err = productDB.Delete(product.ID.String())

	assert.NoError(t, err)
	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}
