package main

import (
    "fmt"
    "time"
)

func main() {
    // Get the current time and print it
    currentTime := time.Now()
    fmt.Println("Current Time:", currentTime)

    // Format and print the current time
    formattedTime := currentTime.Format("2006-01-02 15:04:05")
    fmt.Println("Formatted Time:", formattedTime)
}
// Expected Output: Current Time: 2023-10-06 10:15:30.123456789 +0000 UTC m=+0.000000001
// Formatted Time: 2023-10-06 10:15:30
