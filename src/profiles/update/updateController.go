package profiles

import (
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/profiles/entity"
)

func UpdateController(ctx *gin.Context, profileId string) {
	profile := entity.Profile{}
	err := ctx.BindJSON(&profile)
	if err != nil {
		panic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}

	// service
	res := UpdateService(profile, profileId)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Profile updated successfully", "data": res})
}
