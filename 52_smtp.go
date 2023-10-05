package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func main() {
	// SMTP server configuration
	smtpServer := "smtp.example.com"
	smtpPort := 587
	smtpUsername := "your_username"
	smtpPassword := "your_password"

	// Sender and recipient email addresses
	from := "sender@example.com"
	to := "recipient@example.com"

	// Email message
	subject := "Test Email"
	body := "This is a test email sent using Go."

	// Compose the email message
	msg := "Subject: " + subject + "\n" + body

	// Connect to the SMTP server
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)
	err := smtp.SendMail(smtpServer+":"+fmt.Sprint(smtpPort), auth, from, []string{to}, []byte(msg))
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully.")
}
