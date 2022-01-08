package products

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/products/entity"
)

func ListController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		strLimit := ctx.Query("limit")
		limit, err := strconv.ParseInt(strLimit, 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		listEntity := entity.List{}
		listEntity.Limit = limit
		users := ListService(listEntity)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Products fetched successfully.", "data": users})
	} else { // if error exist
		ctx.JSON(http.StatusBadRequest, gin.H{"data": "Invalid Token"})
	}
}
