package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type ConfigType string

const (
	host     ConfigType = "localhost"
	port     int        = 5432
	user     ConfigType = "postgres"
	password ConfigType = "1"
	dbname   ConfigType = "bookstore"
)

var (
	id         int
	book_name  string
	book_price int
	author_id  int
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\nSuccessfully connected to database!\n")
	}

	// sqlQuery := `INSERT INTO books (id,book_name,book_price,author_id) VALUES($1,$2,$3,$4)`
	// _, err = db.Exec(sqlQuery, 2, "Here we go", 11000, 1)
	//

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &book_name, &book_price, &author_id)
		if err != nil {
			panic(err)
		}
		fmt.Println("\n", id, book_name, book_price, author_id)
	}
	fmt.Printf("\nRow inserted successfully\n")
}
