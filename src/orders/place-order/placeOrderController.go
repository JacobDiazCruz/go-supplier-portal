package orders

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/orders/entity"
	user "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func PlaceOrderController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get email and return user details
		u := user.GetEmail(ct.Username)
		au := entity.Auth{}
		au.UserId = u.ID

		// cart request
		order := entity.PlaceOrder{}
		err := ctx.BindJSON(&order)
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		}
		order.UserId = u.ID

		// update service
		res := PlaceOrderService(order, au)
		fmt.Println(res)

		ctx.JSON(http.StatusOK, gin.H{"msg": "Placed Order Successfully!", "data": res})
	} else { // if error exist
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
