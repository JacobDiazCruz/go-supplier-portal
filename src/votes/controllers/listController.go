package votes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "gitlab.com/JacobDCruz/supplier-portal/src/votes/services"
)

func ListController(ctx *gin.Context) {
	votes := service.ListService()
	ctx.JSON(http.StatusOK, gin.H{"msg": "Votes fetched successfully.", "data": votes})
}
