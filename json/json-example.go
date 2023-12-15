package main
import (
	"fmt"
	"net/http"
	"encoding/json"
)

type User struct {
	FirstName string `json:firstname`
	LastName string `json:lastname`
	Age int `json:age`
}

func encode(w http.ResponseWriter, r *http.Request) {
	john := User{
		FirstName: "John",
		LastName: "Doe",
		Age: 25,
	}
	json.NewEncoder(w).Encode(john)
}

func decode(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Fprintf(w, "%s %s is %d years old!\n", user.FirstName, user.LastName, user.Age)
}

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)

	http.ListenAndServe(":8080", nil)
}
