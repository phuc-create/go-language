package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type ConfigType string

// const (
// 	host     ConfigType = "localhost"
// 	port     int        = 5432
// 	user     ConfigType = "postgres"
// 	password ConfigType = "1"
// 	dbname   ConfigType = "bookstore"
// )

var (
	id         int
	book_name  string
	book_price int
	author_id  int
)
var db *gorm.DB

type Person struct {
	gorm.Model
	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model
	ID       string `json:"name"`
	Author   string `json:"author"`
	Public   bool   `json:"public"`
	PersonID int
}

type Author struct {
	gorm.Model
	ID       string `json:"name"`
	Author   string `json:"author"`
	Public   bool   `json:"public"`
	PersonID int
}

func main() {
	// Connect ------------------------------------------------------------------------------------
	// Get infor connection from .env to avoid sensitive case
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	// create connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)
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

	// Insert Book------------------------------------------------------------------------------------
	POSTGRES_INSERT_STATEMENT := `INSERT INTO books (id,book_name,book_price,author_id) VALUES($1,$2,$3,$4)`
	_, err = db.Exec(POSTGRES_INSERT_STATEMENT)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nRow inserted successfully\n")

	// Update book------------------------------------------------------------------------------------
	POSTGRES_UPDATE_STATEMENT := `UPDATE books
	SET book_name = $1 WHERE id = $2;`
	updateRes, err := db.Exec(POSTGRES_UPDATE_STATEMENT, "Hi Mom!", 2)
	if err != nil {
		panic(err)
	}
	updateRows, err := updateRes.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("\nRow updated ", updateRows)

	// Delete a book------------------------------------------------------------------------------------
	POSTGRES_DELETE_STATEMENT := `DELETE FROM books WHERE id = $1`

	delRes, err := db.Exec(POSTGRES_DELETE_STATEMENT, 4)
	if err != nil {
		panic(err)
	}

	delRow, err := delRes.RowsAffected()
	if err != nil {
		panic(err)
	}

	checkProof := func() string {
		if delRow > 1 {
			return " Row"
		} else {
			return " Rows"
		}
	}

	fmt.Println("\nDelete ", delRow, checkProof())

	// View all books------------------------------------------------------------------------------------
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
}

// expected
// "id"	"book_name"	"book_price"	"author_id" 				// Done
// 1	"H"	11000	1             // stay 								// Done
// 3	"King Kong"	32000	2     // deleted 							// Done
// 2	"Don't give up,Sam"	11000	1 // Update book_name // Done
// 4 "Go Example" 55000 1			// add new rows 				// Done

// "id"	"book_name"	"book_price"	"author_id"
// 3	"King Kong"	32000	2
// 4	"Go Example"	55000	1
// 5	"Do you like Algorithms"	1000000	1
// 2	"Hello World!"	11000	1
