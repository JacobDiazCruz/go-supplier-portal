package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchField struct {
	Search string `json:"search" bson:"search"`
}

func SearchController(ctx *gin.Context) {
	// search := ctx.Query("search")
	search := SearchField{}
	err := ctx.BindJSON(&search)
	if err != nil {
		panic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}
	res := SearchService(search)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Product fetched successfully.", "data": res})
}
