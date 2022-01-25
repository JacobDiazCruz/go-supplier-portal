package users

import (
	"bytes"
	"fmt"
	"net/http"
	"net/smtp"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	authService "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	h "gitlab.com/JacobDCruz/supplier-portal/src/helpers"
)

type ForgotPassword struct {
	Email string `json:"email" bson:"email"`
}

func ForgotController(ctx *gin.Context) {
	fp := ForgotPassword{}
	if err := ctx.ShouldBindJSON(&fp); err == nil {
		validate := validator.New()
		if err := validate.Struct(&fp); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"msg":   h.RequiredField,
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
	}

	// Sender data.
	from := "support@picko.ph"
	password := "aekhrnvfhoffzzum"

	// Receiver email address.
	to := []string{
		fp.Email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	// sign a jwt token
	signToken := authService.SignToken(fp.Email, 10)

	// set body content
	t.Execute(&body, struct {
		Email   string
		Message string
		Link    string
	}{
		Email:   fp.Email,
		Message: "",
		Link:    "http://localhost:3000/reset-password?code=" + signToken + "&email=" + fp.Email,
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Successfully sent, please check your email.",
	})
}
