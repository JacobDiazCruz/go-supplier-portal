package person

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context, userId string) {
	// u := Param{userId}
	// res := []getService{u} // here
	res := GetDataById(userId)
	// fmt.Println(u.id)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Fetched data successfully", "data": res})
}
