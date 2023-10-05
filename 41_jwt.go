package main

import (
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "net/http"
    "time"
)

var jwtKey = []byte("secret_key")

func createTokenHandler(w http.ResponseWriter, r *http.Request) {
    // Create a new token with claims
    claims := jwt.MapClaims{
        "username": "alice",
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        fmt.Println("Error creating token:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    fmt.Fprintln(w, tokenString)
}

func parseTokenHandler(w http.ResponseWriter, r *http.Request) {
    tokenString := r.Header.Get("Authorization")

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Check the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return jwtKey, nil
    })
    if err != nil {
        fmt.Println("Error parsing token:", err)
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        username := claims["username"].(string)
        fmt.Fprintf(w, "Authenticated as %s", username)
    } else {
        fmt.Println("Invalid token")
        w.WriteHeader(http.StatusUnauthorized)
    }
}

func main() {
    http.HandleFunc("/create-token", createTokenHandler)
    http.HandleFunc("/parse-token", parseTokenHandler)
    http.ListenAndServe(":8080", nil)
}
