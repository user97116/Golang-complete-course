package main

import (
    "fmt"
)

func doSomething() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    fmt.Println("Doing something...")
    panic("An error occurred!")
}

func main() {
    // Call a function that uses defer, panic, and recover
    doSomething()
    fmt.Println("Continuing after the function call.")
}
// Expected Output:
// Doing something...
// Recovered: An error occurred!
// Continuing after the function call.
