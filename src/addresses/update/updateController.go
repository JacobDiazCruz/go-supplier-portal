package addresses

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/addresses/entity"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	user "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func UpdateController(ctx *gin.Context, id string) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get email and return user details
		u := user.GetEmail(ct.Email)

		// address request
		address := entity.Address{}
		err := ctx.BindJSON(&address)
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		}
		address.UserId = u.ID

		// service
		res := UpdateService(address, id)
		fmt.Println(res)

		ctx.JSON(http.StatusOK, gin.H{"msg": "Address Updated Successfully!"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
