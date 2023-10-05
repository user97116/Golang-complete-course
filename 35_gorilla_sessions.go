package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("my-secret-key"))

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        session, _ := store.Get(r, "my-session")
        if session.IsNew {
            session.Values["counter"] = 0
        }
        session.Values["counter"] = session.Values["counter"].(int) + 1
        session.Save(r, w)
        fmt.Fprintf(w, "Session Counter: %d", session.Values["counter"])
    })
    http.ListenAndServe(":8080", nil)
}
// Access "http://localhost:8080" in a web browser to increment the session counter.
