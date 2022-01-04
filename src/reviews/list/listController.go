package reviews

import (
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/reviews/entity"
)

func ListController(ctx *gin.Context) {
	productId := ctx.Query("product_id")
	// limit, err := strconv.ParseInt(strLimit, 10, 64)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	listEntity := entity.Review{}
	res := ListService(listEntity, productId)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Reviews fetched successfully.", "data": res})
}
