package db_test

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func setupDb() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	stmt, err := db.Prepare(`create table products(
				"id" string,
				"name" string,
				"price" float,
				"status" string,
				);
`)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	stmt, err := db.Prepare(`insert into products(id, name, price, status) values("1234", "Product 1", 40, "disabled");`)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}
