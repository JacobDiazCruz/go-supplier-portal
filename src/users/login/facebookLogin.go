package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	createCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/create"
	cartEntity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	get "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
	signup "gitlab.com/JacobDCruz/supplier-portal/src/users/signup"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FacebookLogin(ctx *gin.Context) {
	err2 := ctx.BindJSON(&credentials)
	fmt.Println("im here 0")
	if err2 != nil {
		fmt.Println("im here 1")
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}
	// @TODO:
	// VERIFY FB ACCESS TOKEN
	// token := verifyIdToken(credentials.Token)

	// bind requestData
	user := entity.User{
		Email:          credentials.Email,
		Username:       credentials.Username,
		ThumbnailImage: credentials.ThumbnailImage,
		OriginalImage:  credentials.OriginalImage,
		Role:           credentials.Role,
	}

	// validate email if already exist
	emailRes := get.GetEmail(user.Email)

	// Signup email if it doesnt exist yet
	if emailRes.Email == "" {
		res := signup.SignupService(user, "facebook")
		objID, err := primitive.ObjectIDFromHex(res)
		if err != nil {
			panic(err)
		}
		// Create a cart for the new signed up user
		cEntity := &cartEntity.Cart{}
		cEntity.UserId = objID
		cEntity.Products = []string{}
		createCart.AddService(*cEntity)
	}

	// create token
	tk := &auth.TokenIdentity{}
	tk.Email = credentials.Email
	signToken := auth.SignToken(tk.Email, 60)

	ctx.JSON(http.StatusOK, gin.H{"data": signToken})
}
