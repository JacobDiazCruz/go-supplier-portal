package orders

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
)

func GetController(ctx *gin.Context, orderId string) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		res := GetService(orderId)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Fetched data successfully", "data": res})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
