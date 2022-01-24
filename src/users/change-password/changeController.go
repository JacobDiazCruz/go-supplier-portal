package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	getUser "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
	"golang.org/x/crypto/bcrypt"
)

func ChangeController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get email and return user details
		u := getUser.GetEmail(ct.Email)
		// bind requestData
		user := entity.User{}
		err := ctx.BindJSON(&user)
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		}

		// Hashing the password with the default cost of 10
		password := []byte(user.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		user.ID = u.ID
		user.Password = string(hashedPassword)

		// signup service
		res := ChangeService(user)
		fmt.Println(res)

		// http response
		ctx.JSON(http.StatusOK, gin.H{"msg": "Password changed successfully!"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": "Invalid Token"})
	}
}
