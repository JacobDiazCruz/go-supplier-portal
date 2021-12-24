package profiles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListController(ctx *gin.Context) {
	profiles := ListService()
	ctx.JSON(http.StatusOK, gin.H{"msg": "Profiles fetched successfully.", "data": profiles})
}
