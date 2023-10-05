package main

import (
    "encoding/json"
    "fmt"
)

// Define a struct to represent a person
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    // Create a person object
    person := Person{
        Name: "Alice",
        Age:  25,
    }

    // Encode the person object to JSON
    jsonBytes, err := json.Marshal(person)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the JSON representation
    fmt.Println("JSON:", string(jsonBytes))

    // Decode JSON into a person object
    var decodedPerson Person
    err = json.Unmarshal(jsonBytes, &decodedPerson)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the decoded person
    fmt.Printf("Decoded Person: %+v\n", decodedPerson)
}
// Expected Output:
// JSON: {"name":"Alice","age":25}
// Decoded Person: {Name:Alice Age:25}
