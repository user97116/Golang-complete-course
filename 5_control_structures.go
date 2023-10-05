package main

import "fmt"

func main() {
    x := 10

    // Use an if statement to check if x is greater than 5
    if x > 5 {
        fmt.Println("x is greater than 5")
    }

    // Use a for loop to print numbers from 0 to 4
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // Use a switch statement to determine the day of the week
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("It's Monday")
    default:
        fmt.Println("It's not Monday")
    }
}
// Expected Output:
// x is greater than 5
// 0
// 1
// 2
// 3
// 4
// It's Monday
