package main

import (
	"fmt"
	"os"

	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price decimal.Decimal
	gorm.Model
}

func (product *Product) ToString() {
	fmt.Fprintln(os.Stdout, "ID: ", product.ID, " Name: ", product.Name, " Price: ", product.Price)
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	autoMigrate(db)
	create(db)
	createBatch(db)
	firstById(db)
	firstWhere(db)
	findAll(db)
	findPaginated(db)
	findWhere(db)
	save(db)
	delete(db)
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}

func create(db *gorm.DB) {
	db.Create(&Product{
		Name:  "Notebook",
		Price: decimal.NewFromFloat(2999),
	})
}

func createBatch(db *gorm.DB) {
	products := []Product{
		{Name: "Teclado Corsair", Price: decimal.NewFromFloat(599)},
		{Name: "Monitor Dell", Price: decimal.NewFromFloat(1999)},
	}

	db.Create(&products)
}

func firstById(db *gorm.DB) {
	fmt.Println("firstById")
	var product Product
	db.First(&product, 1)
	product.ToString()
}

func firstWhere(db *gorm.DB) {
	fmt.Println("firstWhere")
	var product Product
	db.First(&product, "name LIKE ?", "%teclado%")
	product.ToString()
}

func findAll(db *gorm.DB) {
	fmt.Println("findAll")
	var products []Product
	db.Find(&products)
	for _, p := range products {
		p.ToString()
	}
}

func findPaginated(db *gorm.DB) {
	fmt.Println("findPaginated")
	var products []Product
	db.Limit(2).Offset(2).Find(&products)
	for _, p := range products {
		p.ToString()
	}
}

func findWhere(db *gorm.DB) {
	fmt.Println("findWhere")
	var products []Product
	db.Where("price > ?", 2000).Find(&products)
	for _, p := range products {
		p.ToString()
	}
}

func save(db *gorm.DB) {
	fmt.Println("save")
	var product Product
	db.First(&product, 1)
	product.ToString()

	product.Name = "Mouse"
	db.Save(&product)

	var productUpdated Product
	db.First(&productUpdated, 1)
	productUpdated.ToString()
}

func delete(db *gorm.DB) {
	fmt.Println("delete")
	var product Product
	db.First(&product, 1)
	product.ToString()

	db.Delete(&product)

	findAll(db)
}
