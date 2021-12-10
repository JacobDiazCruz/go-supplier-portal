package profiles

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/profiles/entity"
	service "gitlab.com/JacobDCruz/supplier-portal/src/profiles/services"
)

func UpdateController(ctx *gin.Context, profileId string) {
	fmt.Println("im here update")
	profile := entity.Profile{}
	err := ctx.BindJSON(&profile)
	if err != nil {
		panic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}

	// service
	res := service.UpdateService(profile, profileId)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Profile updated successfully", "data": res})
}
