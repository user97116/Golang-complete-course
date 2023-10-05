package main

import "fmt"

func main() {
    // Create a slice of strings with initial values
    fruits := []string{"apple", "banana"}

    // Append a new element to the slice
    fruits = append(fruits, "cherry")

    // Print the slice
    fmt.Println(fruits)
}
// Expected Output: [apple banana cherry]
