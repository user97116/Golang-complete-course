package main

import (
    "flag"
    "fmt"
)

func main() {
    // Define and parse command-line flags
    var name string
    var age int
    flag.StringVar(&name, "name", "John", "Name of the person")
    flag.IntVar(&age, "age", 30, "Age of the person")
    flag.Parse()

    // Print the values of the flags
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
// Usage: ./program -name Alice -age 25
// Expected Output: Name: Alice, Age: 25
