package addresses

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/addresses/entity"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	user "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func AddController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		// get email and return user details
		u := user.GetEmail(ct.Username)

		// address request
		address := entity.Address{}
		err := ctx.BindJSON(&address)
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		}
		address.UserId = u.ID

		// audit log
		auditLog := &address.AuditLog
		auditLog.Name = ct.Username
		auditLog.Email = ct.Email
		auditLog.ThumbnailImage = ct.ThumbnailImage
		auditLog.OriginalImage = ct.OriginalImage
		auditLog.CreatedAt = time.Now()
		auditLog.CreatedBy = ct.Username
		auditLog.UpdatedAt = time.Now()
		auditLog.UpdatedBy = ct.Username

		// update service
		res := AddService(address)

		// return service
		ctx.JSON(http.StatusOK, gin.H{"msg": "Address added successfully", "data": res})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
