package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "gitlab.com/JacobDCruz/supplier-portal/src/users/services"
)

func GetController(ctx *gin.Context, userId string) {
	res := service.GetService(userId)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Fetched data successfully", "data": res})
}
