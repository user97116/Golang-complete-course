// package1/package1.go
package package1

import "fmt"

func Hello() {
    fmt.Println("Hello from Package1!")
}

// main.go
package main

import (
    "myapp/package1"
)

func main() {
    // Use a custom package in the main program
    package1.Hello()
}
// Expected Output: Hello from Package1!
