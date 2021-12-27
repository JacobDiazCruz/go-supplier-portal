package votes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/votes/entity"
	service "gitlab.com/JacobDCruz/supplier-portal/src/votes/services"
)

func AddController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)
	if ct == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": "Invalid Token"})
		return
	}

	// set audit log
	audit := &entity.AuditLog{
		Name:      ct.Username,
		CreatedAt: time.Now(),
		CreatedBy: ct.Username,
		UpdatedAt: time.Now(),
		UpdatedBy: "",
	}

	// bind request data
	vote := entity.Vote{}
	err2 := ctx.BindJSON(&vote)

	if err2 != nil {
		panic(err2)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		return
	}

	// add service
	queryParams := ctx.Request.URL.Query()
	res := service.AddService(&vote, queryParams["content_id"][0], audit)
	fmt.Println(res)

	// get service details
	getRes := service.GetService(res)
	ctx.JSON(http.StatusOK, gin.H{"data": getRes})
}
