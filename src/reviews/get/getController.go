package reviews

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetController(ctx *gin.Context, id string) {
	res := GetService(id)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Review fetched successfully.", "data": res})
}
