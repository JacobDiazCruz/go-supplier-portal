package carts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListController(ctx *gin.Context) {
	carts := ListService()
	ctx.JSON(http.StatusOK, gin.H{"msg": "Cart Items fetched successfully.", "data": carts})
}
