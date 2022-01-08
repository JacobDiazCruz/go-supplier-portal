package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	get "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func UploadController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get email and return user details
		u := get.GetEmail(ct.Email)

		user := entity.User{}
		err := ctx.BindJSON(&user)
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		}
		user.ID = u.ID

		res := UploadService(user)
		fmt.Println(res)
		// return res
		ctx.JSON(http.StatusOK, gin.H{"msg": "Profile Image uploaded successfully!", "data": ""})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
