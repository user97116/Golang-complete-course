package main

import (
    "fmt"
    "net/http"
    "os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the form data to retrieve the uploaded file
    r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
    file, handler, err := r.FormFile("file")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Create a new file on the server and copy the uploaded file to it
    f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer f.Close()
    _, err = io.Copy(f, file)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Fprintln(w, "File uploaded successfully.")
}

func main() {
    http.HandleFunc("/upload", uploadHandler)
    http.ListenAndServe(":8080", nil)
}
// Use a tool like curl or Postman to make file upload requests to "/upload".
