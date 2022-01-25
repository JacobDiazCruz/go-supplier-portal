package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	createCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/create"
	cartEntity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	h "gitlab.com/JacobDCruz/supplier-portal/src/helpers"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	get "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func SignupController(ctx *gin.Context) {
	// validate user request
	user := entity.User{}
	if err := ctx.ShouldBindJSON(&user); err == nil {
		validate := validator.New()
		if err := validate.Struct(&user); err != nil {
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
	emailRes := get.GetEmail(user.Email)
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

	// signup service
	res := SignupService(user)
	objID, err := primitive.ObjectIDFromHex(res)
	if err != nil {
		panic(err)
	}

	// get user service
	getUser := get.GetService(res)

	// Create a cart for the new signed up user
	cEntity := &cartEntity.Cart{}
	cEntity.UserId = objID
	cEntity.Products = []string{}
	createCart.AddService(*cEntity)

	// http response
	ctx.JSON(http.StatusOK, gin.H{"msg": "Fetched data successfully", "data": getUser})
}
