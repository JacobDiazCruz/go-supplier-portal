package addresses

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/addresses/entity"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	h "gitlab.com/JacobDCruz/supplier-portal/src/helpers"
	user "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func AddController(ctx *gin.Context) {
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
