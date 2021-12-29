package orders

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/orders/entity"
)

func PlaceOrderController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// cart request
		order := entity.PlaceOrder{}
		err := ctx.BindJSON(&order)
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		}

		// update service
		res := PlaceOrderService(order)
		fmt.Println(res)

		ctx.JSON(http.StatusOK, gin.H{"msg": "Placed Order Successfully!", "data": res})
	} else { // if error exist
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
