package votes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "gitlab.com/JacobDCruz/supplier-portal/src/votes/services"
)

func GetController(ctx *gin.Context, voteId string) {
	res := service.GetService(voteId)
	// fmt.Println(u.id)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Fetched data successfully", "data": res})
}
