package products

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/products/entity"
)

func UpdateController(ctx *gin.Context) {
	product := entity.Product{}
	err := ctx.BindJSON(&product)
	if err != nil {
		panic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}

	// service
	res := UpdateService(product)
	fmt.Println(res)

	// getRes := get.GetService(res)
	// fmt.Println(getRes)
	fmt.Println("contents get")
}
