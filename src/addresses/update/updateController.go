package addresses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/addresses/entity"
)

func UpdateController(ctx *gin.Context, profileId string) {
	address := entity.Address{}
	err := ctx.BindJSON(&address)
	if err != nil {
		panic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}

	// service
	res := UpdateService(address, profileId)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Address updated successfully", "data": res})
}
