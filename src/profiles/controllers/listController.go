package profiles

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "gitlab.com/JacobDCruz/supplier-portal/src/profiles/services"
)

func ListController(ctx *gin.Context) {
	profiles := service.ListService()
	ctx.JSON(http.StatusOK, gin.H{"msg": "Profiles fetched successfully.", "data": profiles})
}
