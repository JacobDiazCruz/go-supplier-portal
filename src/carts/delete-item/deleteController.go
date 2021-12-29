package carts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
)

func DeleteController(ctx *gin.Context, id string) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		cart := entity.ProductRequest{}
		err := ctx.BindJSON(&cart)
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		}

		// update service
		res := DeleteService(cart, id)
		fmt.Println(res)

		ctx.JSON(http.StatusOK, gin.H{"msg": "Cart Item Removed Successfully"})
	} else { // if error exist
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
