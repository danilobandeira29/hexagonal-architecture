package main

import (
	"database/sql"
	db2 "github.com/danilobandeira29/hexagonal-architecture/adapter/db"
	"github.com/danilobandeira29/hexagonal-architecture/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product example", 44.0)
	productService.Enable(product)
}
