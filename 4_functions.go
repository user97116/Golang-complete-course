package main

import "fmt"

// Define a function that adds two integers and returns the result
func add(a, b int) int {
    return a + b
}

func main() {
    // Call the add function and print the result
    result := add(3, 5)
    fmt.Printf("Result of addition: %d\n", result)
}
// Expected Output: Result of addition: 8
