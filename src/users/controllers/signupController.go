package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	service "gitlab.com/JacobDCruz/supplier-portal/src/users/services"
	"golang.org/x/crypto/bcrypt"
)

func SignupController(ctx *gin.Context) string {
	user := entity.User{}
	err2 := ctx.BindJSON(&user)
	if err2 != nil {
		panic(err2)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}

	// Hashing the password with the default cost of 10
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)

	// err = validate.Struct(person)
	// if err != nil {
	// 	return err
	// }
	res := service.SignupService(user)
	fmt.Println(res)
	return res
}
