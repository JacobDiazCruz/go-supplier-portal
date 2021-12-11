package votes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/votes/entity"
	service "gitlab.com/JacobDCruz/supplier-portal/src/votes/services"
)

func AddController(ctx *gin.Context) {
	// check token and return
	ct, err := auth.GetToken(ctx)
	fmt.Println(ct)
	fmt.Println("qwrqwrqwrqrqr")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": "Invalid Token"})
		return
	}
	if ct == "Bad Request" || ct == "Unauthorized" {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": "Invalid Token"})
		return
	}

	// bind request data
	vote := entity.Vote{}
	err2 := ctx.BindJSON(&vote)
	if err2 != nil {
		panic(err2)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		return
	}

	// service
	queryParams := ctx.Request.URL.Query()
	res := service.AddService(vote, queryParams["content_id"][0])
	fmt.Println(res)

	// get service details
	getRes := service.GetService(res)
	ctx.JSON(http.StatusOK, gin.H{"data": getRes})
}
