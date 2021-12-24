package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	addProfile "gitlab.com/JacobDCruz/supplier-portal/src/profiles/add"
	profilesEntity "gitlab.com/JacobDCruz/supplier-portal/src/profiles/entity"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	get "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func SignupController(ctx *gin.Context) {
	// bind requestData
	user := entity.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		panic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}

	// validate email if already exist
	emailRes, err := get.GetEmail(user.Email)
	if err != nil {
		fmt.Println("Err")
	}
	if emailRes.Email != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Email already exist."})
		return
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

	// signup service
	res := SignupService(user)
	objID, err := primitive.ObjectIDFromHex(res)
	if err != nil {
		panic(err)
	}

	// get user service
	getUser := get.GetService(res)

	// Create profile
	pEntity := &profilesEntity.Profile{
		UserId:    objID,
		Email:     getUser.Email,
		FirstName: getUser.FirstName,
		LastName:  getUser.LastName,
		Role:      getUser.Role,
	}
	addProfile.AddService(*pEntity)

	// http response
	ctx.JSON(http.StatusOK, gin.H{"msg": "Fetched data successfully", "data": getUser})
}
