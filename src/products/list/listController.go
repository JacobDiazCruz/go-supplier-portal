package products

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/products/entity"
)

func ListController(ctx *gin.Context) {
	strLimit := ctx.Query("limit")

	// validate query params
	if strLimit == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error fetching product. No limit parameter.", "data": ""})
	}

	// service
	limit, err := strconv.ParseInt(strLimit, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	listEntity := entity.List{}
	listEntity.Limit = limit
	users := ListService(listEntity)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Products fetched successfully.", "data": users})
}
