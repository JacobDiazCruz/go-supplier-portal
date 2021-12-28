package carts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
)

func AddController(ctx *gin.Context) {
	cart := entity.Cart{}
	err := ctx.BindJSON(&cart)
	if err != nil {
		panic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}

	// service
	res := AddService(cart)
	fmt.Println(res)

	// get cart
	fmt.Println("carts get")
}
