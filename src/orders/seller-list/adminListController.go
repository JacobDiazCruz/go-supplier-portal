package orders

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	user "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func ListController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get email and return user details
		u := user.GetEmail(ct.Username)

		// get service and return json
		res := ListService(u.ID)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Orders fetched successfully.", "data": res})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
