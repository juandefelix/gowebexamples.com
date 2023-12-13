package main

import (
	"fmt"
	"time"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://user:password@localhost/test_database"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	type user struct {
		id int
		firstName string
		lastName string
		email string
		createdAt time.Time
	}


	{ // Query all users
		rows, err := db.Query(`SELECT id, first_name, last_name, email, created_at FROM admin_users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user
			err = rows.Scan(&u.id, &u.firstName, &u.lastName, &u.email, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}

		fmt.Printf("%#v", users)

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
