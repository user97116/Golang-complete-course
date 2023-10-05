package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{"Alice", 30}

    // Get the type of p and print its fields
    t := reflect.TypeOf(p)
    fmt.Println("Type:", t)
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Field %d: %s (%s)\n", i, field.Name, field.Type)
    }

    // Get the value of p and print its field values
    v := reflect.ValueOf(p)
    fmt.Println("Value:", v)
    for i := 0; i < t.NumField(); i++ {
        field := v.Field(i)
        fmt.Printf("Field %d: %v\n", i, field.Interface())
    }
}
// Expected Output:
// Type: main.Person
// Field 0: Name (string)
// Field 1: Age (int)
// Value: {Alice 30}
// Field 0: Alice
// Field 1: 30
