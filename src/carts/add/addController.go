package carts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	get "gitlab.com/JacobDCruz/supplier-portal/src/carts/get"
	user "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func AddController(ctx *gin.Context, id string) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get email and return user details
		u := user.GetEmail(ct.Username)
		fmt.Println("here user123123123123123")

		// cart request
		cart := entity.ProductRequest{}
		err := ctx.BindJSON(&cart)
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		}
		cart.UserId = u.ID
		// cart.ProductId = m.Hex()

		// update service
		res := AddService(cart, id)
		fmt.Println(res)

		// get details and return json
		getRes := get.GetService(res)
		fmt.Println(getRes)
		fmt.Println("contents get")
		ctx.JSON(http.StatusOK, gin.H{"msg": "Cart Item Updated Successfully", "data": getRes})
	} else { // if error exist
		ctx.JSON(http.StatusBadRequest, gin.H{"data": "Invalid Token"})
	}
}
