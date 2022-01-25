package users

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	h "gitlab.com/JacobDCruz/supplier-portal/src/helpers"
	get "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

type VerifyEmail struct {
	Email string `json:"email" bson:"email"`
}

func VerifyEmailController(ctx *gin.Context) {
	// validate user request
	ve := VerifyEmail{}
	if err := ctx.ShouldBindJSON(&ve); err == nil {
		validate := validator.New()
		if err := validate.Struct(&ve); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"msg":   h.RequiredField,
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
	}

	// validate email if already exist
	emailRes := get.GetEmail(ve.Email)
	if emailRes.Email != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Email already exist."})
		return
	}

	// signup service
	// store code on db and send it to email
	code := MakeVerificationCode()
	se := SendEmailVerification(ctx, ve, code)

	// http response
	ctx.JSON(http.StatusOK, gin.H{"msg": se})
}

func VerifyCodeController(ctx *gin.Context) {
	// validate user request
	codeStruct := VerifyCode{}
	if err := ctx.ShouldBindJSON(&codeStruct); err == nil {
		validate := validator.New()
		if err := validate.Struct(&codeStruct); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"msg":   h.RequiredField,
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
	}

	// @TODO: Verify code if it exist in db
	verifiedCode := GetVerificationCode(codeStruct.Code)
	fmt.Println(verifiedCode)

	// http response
	ctx.JSON(http.StatusOK, gin.H{"msg": "Email verified successfully!"})
}

func SendEmailVerification(ctx *gin.Context, ve VerifyEmail, code string) string {

	// Sender data.
	from := "support@picko.ph"
	password := "aekhrnvfhoffzzum"

	// Receiver email address.
	to := []string{
		ve.Email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("email_verification.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	// set body content
	t.Execute(&body, struct {
		Email   string
		Message string
		Code    string
	}{
		Email: ve.Email,
		Code:  code,
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return "Error sending email"
	}
	fmt.Println("Email Sent!")
	return "Email sent! Please verify your email."
}
