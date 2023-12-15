package main
import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	fmt.Println(session.Values)

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	// Secret message
	fmt.Fprintf(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	// Authentication goes here

	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/secret", secret)

	http.ListenAndServe(":8080", nil)
}
