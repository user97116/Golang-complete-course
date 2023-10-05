package main

import "fmt"

func main() {
    // Define a map to store a person's age
    person := map[string]int{
        "age": 30,
    }

    // Retrieve and print the age from the map
    fmt.Println(person["age"])
}
// Expected Output: 30
