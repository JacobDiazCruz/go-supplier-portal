package carts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	user "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func ValidateController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get email and return user details
		u := user.GetEmail(ct.Email)

		// service
		userId := u.ID
		productsToUpdate := ValidateService(userId)
		if len(productsToUpdate) > 0 {
			ctx.JSON(http.StatusOK, gin.H{"msg": "You have outdated products in your cart.", "data": productsToUpdate})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"msg": "Your cart items are valid.", "data": productsToUpdate})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": "Invalid Token"})
	}
}
