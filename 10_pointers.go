package main

import "fmt"

func main() {
    // Declare a variable and a pointer to it
    x := 10
    y := &x

    // Print the value of x and the memory address of y
    fmt.Printf("x: %d, y: %p\n", x, y)
}
// Expected Output: x: 10, y: 0xc00008a008
