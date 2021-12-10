package contents

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "gitlab.com/JacobDCruz/supplier-portal/src/contents/services"
)

func GetController(ctx *gin.Context, userId string) {
	res := service.GetService(userId)
	// fmt.Println(u.id)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Fetched data successfully", "data": res})
}
