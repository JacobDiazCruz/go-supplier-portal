package carts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetController(ctx *gin.Context, id string) {
	carts := GetService(id)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Cart Items fetched successfully.", "data": carts})
}
