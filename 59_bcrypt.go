package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "mysecretpassword"

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Hashed Password:", string(hashedPassword))

	// Verify a password
	inputPassword := "wrongpassword"
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(inputPassword))
	if err == nil {
		fmt.Println("Password is correct.")
	} else {
		fmt.Println("Password is incorrect.")
	}
}
