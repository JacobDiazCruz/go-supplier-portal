package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
)

func LoginController(ctx *gin.Context) {
	credentials := entity.Credentials{}

	err := ctx.BindJSON(&credentials)

	// if payload is invalid
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Bad Request"})
		return
	}

	res := LoginService(&credentials)

	ctx.SetCookie("token", res, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}
