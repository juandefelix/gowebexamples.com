package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://juanortiz@localhost/standard_communities_dev?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecing to the database:", err)
	}

	rows, err := db.Query("SELECT email FROM admin_users")
	if err != nil {
		fmt.Println("Error reading database:", err)
	}

	fmt.Println("Results:")
	fmt.Println(rows.Columns())
}
