package products

import (
	"fmt"
	"net/http"

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
