package person

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/JacobDCruz/supplier-portal/src/auth"
)

func DeleteController(ctx *gin.Context, id string) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		DeleteService(id)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Data success"})
	} else { // if error exist
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
