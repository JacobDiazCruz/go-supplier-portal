package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	profilesEntity "gitlab.com/JacobDCruz/supplier-portal/src/profiles/entity"
	profilesService "gitlab.com/JacobDCruz/supplier-portal/src/profiles/services"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	service "gitlab.com/JacobDCruz/supplier-portal/src/users/services"
	"golang.org/x/crypto/bcrypt"
)

func SignupController(ctx *gin.Context) {
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
	user.Password = string(hashedPassword)

	// err = validate.Struct(person)
	// if err != nil {
	// 	return err
	// }

	res := service.SignupService(user)
	getUser := service.GetService(res)

	// Create profile
	pEntity := &profilesEntity.Profile{
		UserId:    res,
		Email:     getUser.Email,
		FirstName: getUser.FirstName,
		LastName:  getUser.LastName,
		Role:      getUser.Role,
	}
	profileRes := profilesService.AddService(*pEntity)
	fmt.Println(profileRes)
	fmt.Println("profileRes")
}
