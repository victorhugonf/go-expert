package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID    string
	Name  string
	Price decimal.Decimal
}

func NewProduct(name string, price decimal.Decimal) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = createTableProducts(db)
	if err != nil {
		panic(err)
	}
	product := NewProduct("Notebook", decimal.NewFromFloat32(1999.99))
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Price = decimal.NewFromFloat32(2199.99)
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}
	p, err := findProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	price, _ := p.Price.Float64()
	fmt.Println("Find by id:")
	fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, price)
	fmt.Println("find all products:")
	products, err := findAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		price, _ := p.Price.Float64()
		fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, price)
	}
	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

func createTableProducts(db *sql.DB) error {
	_, err := db.Exec("create table if not exists products(id varchar(255), name varchar(80), price decimal, primary key(id))")
	if err != nil {
		return err
	}
	return nil
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func findProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func findAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
