package users

import (
	"fmt"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

func ForgotController(ctx *gin.Context) {
	// Sender data.
	from := "carljacobdiazcruz@gmail.com"
	password := "Ceejay33"

	// Receiver email address.
	to := []string{
		"joojojo@yopmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
