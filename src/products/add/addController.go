package products

import (
	"crypto/rand"
	"encoding/hex"
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
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered in adding a product"})
		}

		// product audit log
		auditLog := &product.AuditLog
		auditLog.Name = ct.Username
		auditLog.Email = ct.Email
		auditLog.ThumbnailImage = ct.ThumbnailImage
		auditLog.OriginalImage = ct.OriginalImage
		auditLog.CreatedAt = time.Now()
		auditLog.CreatedBy = ct.Username
		auditLog.UpdatedAt = time.Now()
		auditLog.UpdatedBy = ct.Username

		// ======================================
		// Approach 1 Variant option collection:
		// Problem: slow if we create all variant options in 1 api
		variants := product.Variants
		for variantKey, variant := range variants {
			b := make([]byte, 10) // equals 16 characters
			rand.Read(b)
			randId := hex.EncodeToString(b)
			variants[variantKey].ID = randId
			for optionKey, _ := range variant.Options {
				b := make([]byte, 10) // equals 16 characters
				rand.Read(b)
				randId := hex.EncodeToString(b)
				variants[variantKey].Options[optionKey].ID = randId
			}
		}

		// service
		res := AddProductService(product)
		fmt.Println(res)

		// get product
		getRes := get.GetService(res, "")
		ctx.JSON(http.StatusOK, gin.H{"msg": "Product added successfully.", "data": getRes})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
