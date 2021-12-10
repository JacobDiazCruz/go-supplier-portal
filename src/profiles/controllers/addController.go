package profiles

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/profiles/entity"
	service "gitlab.com/JacobDCruz/supplier-portal/src/profiles/services"
)

func AddController(ctx *gin.Context) string {
	profile := entity.Profile{}
	err := ctx.BindJSON(&profile)
	if err != nil {
		panic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}

	// service
	res := service.AddService(profile)
	fmt.Println(res)
	return res
}
