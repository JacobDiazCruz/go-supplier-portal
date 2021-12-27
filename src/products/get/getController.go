package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetController(ctx *gin.Context) {
	productId := ctx.Query("id")
	slug := ctx.Query("slug")
	// param := Param{
	// 	id:   productId,
	// 	slug: slug,
	// }
	res := GetService(productId, slug)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Product fetched successfully.", "data": res})
}
