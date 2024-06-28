package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Category HAS MANY Product
type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

// Product BELONGS TO Category
// Product HAS ONE SerialNumber
type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        decimal.Decimal
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func (product *Product) ToString() {
	fmt.Fprintln(os.Stdout,
		"ID:", product.ID,
		"Name:", product.Name,
		"Price:", product.Price,
		"Category ID:", product.Category.ID,
		"Category name:", product.Category.Name,
		"Serial number ID:", product.SerialNumber.ID,
		"Serial number:", product.SerialNumber.Number)
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	autoMigrate(db)
	create(db)
	findAllProducts(db)
	findAllCategoriesAndProducts(db)
	findAllCategoriesAndProductsAndSerialNumber(db)
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})
}

func create(db *gorm.DB) {
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	product := Product{
		Name:       "Notebook",
		Price:      decimal.NewFromFloat(2999),
		CategoryID: category.ID,
	}
	db.Create(&product)

	db.Create(&SerialNumber{
		Number:    uuid.NewString(),
		ProductID: product.ID,
	})

	product = Product{
		Name:       "Teclado",
		Price:      decimal.NewFromFloat(499),
		CategoryID: category.ID,
	}
	db.Create(&product)

	db.Create(&SerialNumber{
		Number:    uuid.NewString(),
		ProductID: product.ID,
	})
}

func findAllProducts(db *gorm.DB) {
	fmt.Println("findAllProducts")
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, p := range products {
		p.ToString()
	}
}

func findAllCategoriesAndProducts(db *gorm.DB) {
	fmt.Println("findAllCategoriesAndProducts")
	var categories []Category
	err := db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			fmt.Println("-", p.Name, p.Category.Name)
		}
	}
}

func findAllCategoriesAndProductsAndSerialNumber(db *gorm.DB) {
	fmt.Println("findAllCategoriesAndProductsAndSerialNumber")
	var categories []Category
	//err := db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	//OU
	err := db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			fmt.Println("-", p.Name, p.Category.Name, p.SerialNumber.Number)
		}
	}
}
