package addresses

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/addresses/entity"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	h "gitlab.com/JacobDCruz/supplier-portal/src/helpers"
	user "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func UpdateController(ctx *gin.Context, id string) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get email and return user details
		u := user.GetEmail(ct.Email)
		address := entity.Address{}

		// address request
		if err := ctx.ShouldBindJSON(&address); err == nil {
			validate := validator.New()
			if err := validate.Struct(&address); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code":  http.StatusBadRequest,
					"msg":   h.RequiredField,
					"error": err.Error(),
				})
				ctx.Abort()
				return
			}
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
