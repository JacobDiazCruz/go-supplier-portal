package carts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	user "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func ClearController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get email and return user details
		u := user.GetEmail(ct.Username)

		// cart request
		cart := entity.ProductRequest{}
		cart.UserId = u.ID

		// update service
		res := ClearService(cart)
		fmt.Println(res)

		ctx.JSON(http.StatusOK, gin.H{"msg": "Cart Item Removed Successfully"})
	} else { // if error exist
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
