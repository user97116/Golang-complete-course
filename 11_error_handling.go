package main

import (
    "errors"
    "fmt"
)

// CustomError represents a custom error type
type CustomError struct {
    message string
}

func (e *CustomError) Error() string {
    return e.message
}

func divide(a, b int) (int, error) {
    // Check for division by zero and return a custom error if necessary
    if b == 0 {
        return 0, &CustomError{"division by zero"}
    }
    return a / b, nil
}

func main() {
    // Perform division and handle errors
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
// Expected Output: Result: 5
