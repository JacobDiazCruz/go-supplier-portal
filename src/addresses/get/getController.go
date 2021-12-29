package addresses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetController(ctx *gin.Context) {
	profileId := ctx.Query("profile_id")
	res := GetService(profileId)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Product fetched successfully.", "data": res})
}
