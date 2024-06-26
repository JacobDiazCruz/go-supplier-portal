package addresses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
)

func GetController(ctx *gin.Context, id string) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get service and return json
		res := GetService(id)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Product fetched successfully.", "data": res})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
