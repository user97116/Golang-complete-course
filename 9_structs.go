package main

import "fmt"

// Define a struct to represent a person with name and age
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create an instance of the Person struct
    person := Person{
        Name: "Alice",
        Age:  25,
    }

    // Print the person's name and age
    fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
// Expected Output: Name: Alice, Age: 25
