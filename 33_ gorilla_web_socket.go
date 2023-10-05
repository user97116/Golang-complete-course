package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
    // Upgrade the HTTP connection to a WebSocket connection
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()

    // Echo messages received from the WebSocket client
    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        if err := conn.WriteMessage(messageType, p); err != nil {
            fmt.Println("Error:", err)
            return
        }
    }
}

func main() {
    http.HandleFunc("/ws", handler)
    http.ListenAndServe(":8080", nil)
}
