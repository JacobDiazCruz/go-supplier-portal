package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetController(ctx *gin.Context) {
	productId := ctx.Query("id")
	slug := ctx.Query("slug")

	// validate
	if slug == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error fetching product. Invalid slug.", "data": ""})
	}

	// service
	res := GetService(productId, slug)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Product fetched successfully.", "data": res})
}
