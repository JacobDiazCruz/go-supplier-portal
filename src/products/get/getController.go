package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetController(ctx *gin.Context) {
	userId := ctx.Query("id")
	slug := ctx.Query("slug")
	param := Param{
		id:   userId,
		slug: slug,
	}
	res := getService.GetService(param)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Products fetched successfully.", "data": res})
}
