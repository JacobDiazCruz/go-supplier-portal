package contents

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "gitlab.com/JacobDCruz/supplier-portal/src/contents/services"
)

func ListController(ctx *gin.Context) {
	contents := service.ListService()
	ctx.JSON(http.StatusOK, gin.H{"msg": "Contents fetched successfully.", "data": contents})
}
