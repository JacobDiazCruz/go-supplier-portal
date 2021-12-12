package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	profilesEntity "gitlab.com/JacobDCruz/supplier-portal/src/profiles/entity"
	profilesService "gitlab.com/JacobDCruz/supplier-portal/src/profiles/services"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	service "gitlab.com/JacobDCruz/supplier-portal/src/users/services"
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
	emailRes, err := service.GetEmail(user.Email)
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
	res := service.SignupService(user)
	objID, err := primitive.ObjectIDFromHex(res)
	if err != nil {
		panic(err)
	}

	// get user service
	getUser := service.GetService(res)

	// Create profile
	pEntity := &profilesEntity.Profile{
		UserId:    objID,
		Email:     getUser.Email,
		FirstName: getUser.FirstName,
		LastName:  getUser.LastName,
		Role:      getUser.Role,
	}
	profilesService.AddService(*pEntity)

	// http response
	ctx.JSON(http.StatusOK, gin.H{"msg": "Fetched data successfully", "data": getUser})
}
