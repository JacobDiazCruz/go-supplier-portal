package products

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/products/entity"
	get "gitlab.com/JacobDCruz/supplier-portal/src/products/get"
)

func AddController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		product := entity.Product{}
		err := ctx.BindJSON(&product)
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		}

		// audit log
		auditLog := &product.AuditLog
		auditLog.Name = ct.Username
		auditLog.Email = ct.Email
		auditLog.ThumbnailImage = ct.ThumbnailImage
		auditLog.OriginalImage = ct.OriginalImage
		auditLog.CreatedAt = time.Now()
		auditLog.CreatedBy = ct.Username
		auditLog.UpdatedAt = time.Now()
		auditLog.UpdatedBy = ct.Username

		// service
		res := AddService(product)
		fmt.Println(res)

		// get product
		getRes := get.GetService(res, "")
		ctx.JSON(http.StatusOK, gin.H{"msg": "Product added successfully.", "data": getRes})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
