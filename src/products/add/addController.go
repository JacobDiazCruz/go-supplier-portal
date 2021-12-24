package products

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/products/entity"
)

func AddController(ctx *gin.Context) {
	product := entity.Product{}
	err := ctx.BindJSON(&product)
	if err != nil {
		panic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}

	// service
	res := AddService(product)
	fmt.Println(res)

	// getRes := get.GetService(res)
	// fmt.Println(getRes)
	fmt.Println("contents get")
}

// func New(g get.getService) {

// }
