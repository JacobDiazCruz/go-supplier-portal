package contents

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/contents/entity"
	service "gitlab.com/JacobDCruz/supplier-portal/src/contents/services"
)

func AddController(ctx *gin.Context) {
	content := entity.Content{}
	err := ctx.BindJSON(&content)
	if err != nil {
		panic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}

	// service
	res := service.AddService(content)
	fmt.Println(res)

	getRes := service.GetService(res)
	fmt.Println(getRes)
	fmt.Println("contents get")
}
