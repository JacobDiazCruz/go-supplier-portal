package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchController(ctx *gin.Context) {
	search := ctx.Query("search")
	res := SearchService(search)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Product fetched successfully.", "data": res})
}
