package person

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/person/entity"
)

func AddUser(ctx *gin.Context) string {
	person := entity.Employee{}
	err := ctx.BindJSON(&person)
	if err != nil {
		panic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}
	// err = validate.Struct(person)
	// if err != nil {
	// 	return err
	// }
	res := SaveUser(person)
	fmt.Println(res)
	return res
	// ctx.JSON(http.StatusOK, gin.H{"msg": "Data success", "data": res})
}
