package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListController(ctx *gin.Context) {
	users := ListService()
	ctx.JSON(http.StatusOK, gin.H{"msg": "Users fetched successfully.", "data": users})
}
