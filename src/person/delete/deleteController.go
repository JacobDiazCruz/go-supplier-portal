package person

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteController(ctx *gin.Context, id string) {
	// if id != nil {
	// 	panic(err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Id is invalid"})
	// }
	res := DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Data success", "data": res})
}
