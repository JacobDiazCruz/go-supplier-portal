package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetController(ctx *gin.Context, email string) {
	res := GetEmail(email)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Fetched data successfully", "data": res})
}
