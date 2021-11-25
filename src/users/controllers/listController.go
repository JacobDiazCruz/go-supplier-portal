package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "gitlab.com/JacobDCruz/supplier-portal/src/users/services"
)

func ListController(ctx *gin.Context) {
	users := service.ListService()
	ctx.JSON(http.StatusOK, gin.H{"msg": "Users fetched successfully.", "data": users})
}
