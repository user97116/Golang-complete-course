package main

import (
    "fmt"
    "net/http"
)

type Greeter interface {
    Greet() string
}

type EnglishGreeter struct{}

func (eg EnglishGreeter) Greet() string {
    return "Hello"
}

type SpanishGreeter struct{}

func (sg SpanishGreeter) Greet() string {
    return "Hola"
}

func greetHandler(w http.ResponseWriter, r *http.Request, greeter Greeter) {
    greeting := greeter.Greet()
    fmt.Fprintln(w, greeting)
}

func main() {
    engGreeter := EnglishGreeter{}
    spaGreeter := SpanishGreeter{}

    http.HandleFunc("/english", func(w http.ResponseWriter, r *http.Request) {
        greetHandler(w, r, engGreeter)
    })

    http.HandleFunc("/spanish", func(w http.ResponseWriter, r *http.Request) {
        greetHandler(w, r, spaGreeter)
    })

    http.ListenAndServe(":8080", nil)
}
