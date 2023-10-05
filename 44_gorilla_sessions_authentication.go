package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("my-secret-key"))

func loginHandler(w http.ResponseWriter, r *http.Request) {
    // Simulate user authentication
    username := "alice"
    session, _ := store.Get(r, "my-session")
    session.Values["username"] = username
    session.Save(r, w)
    fmt.Fprintln(w, "Logged in as", username)
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "my-session")
    if username, ok := session.Values["username"].(string); ok {
        fmt.Fprintln(w, "Protected resource for", username)
    } else {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
    }
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/login", loginHandler)
    r.HandleFunc("/protected", protectedHandler)
    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
// Access "/login" to log in, and then access "/protected" to see the protected resource.
