package carts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	user "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func GetController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get email and return user details
		u := user.GetEmail(ct.Username)

		// service
		userId := u.ID
		carts := GetService(userId)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Cart Items fetched successfully.", "data": carts})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": "Invalid Token"})
	}
}
