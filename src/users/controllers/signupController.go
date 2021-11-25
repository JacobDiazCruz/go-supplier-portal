package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	service "gitlab.com/JacobDCruz/supplier-portal/src/users/services"
)

func SignupController(ctx *gin.Context) string {
	person := entity.User{}
	err := ctx.BindJSON(&person)
	if err != nil {
		panic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}
	// err = validate.Struct(person)
	// if err != nil {
	// 	return err
	// }
	res := service.SignupService(person)
	fmt.Println(res)
	return res
	// ctx.JSON(http.StatusOK, gin.H{"msg": "Data success", "data": res})
}
