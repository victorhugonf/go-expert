package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      decimal.Decimal
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func (c *Category) ToString() {
	fmt.Fprintln(os.Stdout,
		"Name:", c.Name,
	)
}

func (p *Product) ToString() {
	var categories []string
	for _, c := range p.Categories {
		categories = append(categories, c.Name)
	}
	fmt.Fprintln(os.Stdout,
		"ID:", p.ID,
		"Name:", p.Name,
		"Price:", p.Price,
		"Categories:", strings.Join(categories, ", "),
	)
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
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Product{}, &Category{})
}

func create(db *gorm.DB) {
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	category2 := Category{Name: "Cozinha"}
	db.Create(&category2)

	db.Create(&Product{
		Name:       "Panela eletrica",
		Price:      decimal.NewFromFloat(399),
		Categories: []Category{category, category2},
	})

	db.Create(&Product{
		Name:       "Sanduicheira",
		Price:      decimal.NewFromFloat(99),
		Categories: []Category{category, category2},
	})
}

func findAllProducts(db *gorm.DB) {
	fmt.Println("findAllProducts")
	var products []Product
	db.Preload("Categories").Find(&products)
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
			fmt.Println("-", p.Name)
		}
	}
}
