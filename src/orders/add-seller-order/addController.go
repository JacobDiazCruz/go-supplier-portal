package orders

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/orders/entity"
)

func AddController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		order := entity.SellerOrder{}
		err := ctx.BindJSON(&order)
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		}

		// audit log
		auditLog := &order.AuditLog
		auditLog.Name = ct.Username
		auditLog.Email = ct.Email
		auditLog.ThumbnailImage = ct.ThumbnailImage
		auditLog.OriginalImage = ct.OriginalImage
		auditLog.CreatedAt = time.Now()
		auditLog.CreatedBy = ct.Username
		auditLog.UpdatedAt = time.Now()
		auditLog.UpdatedBy = ct.Username

		// service
		res := AddService(order)
		fmt.Println(res)

		// get product
		ctx.JSON(http.StatusOK, gin.H{"msg": "Order added successfully.", "data": ""})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
