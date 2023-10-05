package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Connect to a MySQL database
    db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/mydb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Execute a simple SQL query
    rows, err := db.Query("SELECT name FROM users WHERE id = ?", 1)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var name string
    for rows.Next() {
        err := rows.Scan(&name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Name:", name)
    }
}
// Expected Output: Name: John
