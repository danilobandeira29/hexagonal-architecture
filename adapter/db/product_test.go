package db_test

import (
	"database/sql"
	"github.com/danilobandeira29/hexagonal-architecture/adapter/db"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
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
				"status" string
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

func TestProductDb_Get(t *testing.T) {
	setupDb()
	productDb := db.NewProductDb(Db)
	defer Db.Close()
	product, err := productDb.Get("1234")
	require.Nil(t, err)
	require.Equal(t, "1234", product.GetID())
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 40.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}
