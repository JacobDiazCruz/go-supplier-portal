package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	h "gitlab.com/JacobDCruz/supplier-portal/src/helpers"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
)

func LoginController(ctx *gin.Context) {
	credentials := entity.Credentials{}

	// validate user request
	if err := ctx.ShouldBindJSON(&credentials); err == nil {
		validate := validator.New()
		if err := validate.Struct(&credentials); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"msg":   h.RequiredField,
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
	}

	res := LoginService(&credentials)
	if res == "Error" {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid email or password", "data": ""})
	} else {
		ctx.SetCookie("token", res, 3600, "/", "localhost", false, true)
		ctx.JSON(http.StatusOK, gin.H{"msg": "User logged in successfully!", "data": res})
	}
}
