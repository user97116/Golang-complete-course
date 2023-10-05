package main

import (
    "fmt"
    "os"
)

func main() {
    // Retrieve and print command line arguments
    args := os.Args
    fmt.Println("Command Line Arguments:", args)
}
// Expected Output: Command Line Arguments: [executable_name arg1 arg2 ...]
