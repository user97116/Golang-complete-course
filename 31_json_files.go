package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

// Define a struct to represent a configuration
type Config struct {
    Name  string `json:"name"`
    Value int    `json:"value"`
}

func main() {
    // Read a JSON file and decode it into a Config struct
    fileContents, err := ioutil.ReadFile("config.json")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    var config Config
    err = json.Unmarshal(fileContents, &config)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Config Name: %s, Value: %d\n", config.Name, config.Value)
}
// Expected Output: Config Name: Example Config, Value: 42
