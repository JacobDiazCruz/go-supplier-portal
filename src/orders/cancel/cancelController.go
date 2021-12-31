package orders

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
)

func CancelController(ctx *gin.Context, orderId string) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get email and return user details
		// u := user.GetEmail(ct.Username)

		// cart request
		// cart := entity.ProductRequest{}
		// err := ctx.BindJSON(&cart)
		// if err != nil {
		// 	panic(err)
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		// }
		// cart.UserId = u.ID
		// cart.ProductId = productId

		// update service
		res := CancelService(orderId)
		fmt.Println(res)

		// get details and return json
		// getRes := get.GetService(res)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Order Canceled"})
	} else { // if error exist
		ctx.JSON(http.StatusBadRequest, gin.H{"data": "Invalid Token"})
	}
}
