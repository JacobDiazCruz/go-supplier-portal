package carts

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	user "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func AddController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get email and return user details
		u := user.GetEmail(ct.Email)

		// cart request
		cart := entity.AddToCart{}
		err := ctx.BindJSON(&cart)
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		}
		cart.UserId = u.ID

		// audit log
		auditLog := &cart.AuditLog
		auditLog.Name = ct.Username
		auditLog.Email = ct.Email
		auditLog.ThumbnailImage = ct.ThumbnailImage
		auditLog.OriginalImage = ct.OriginalImage
		auditLog.CreatedAt = time.Now()
		auditLog.CreatedBy = ct.Username
		auditLog.UpdatedAt = time.Now()
		auditLog.UpdatedBy = ct.Username

		// update service
		res := AddService(cart)
		fmt.Println(res)

		ctx.JSON(http.StatusOK, gin.H{"msg": "Cart Item Added Successfully"})
	} else { // if error exist
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
