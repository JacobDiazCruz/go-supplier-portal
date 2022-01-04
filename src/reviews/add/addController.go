package reviews

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/reviews/entity"
)

func AddController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	// if no error
	if ct != nil {
		review := entity.Review{}
		err := ctx.BindJSON(&review)
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		}
		// err = validate.Struct(person)
		// if err != nil {
		// 	return err
		// }
		res := AddService(review)
		fmt.Println(res)
		// return res
		ctx.JSON(http.StatusOK, gin.H{"msg": "Product added successfully.", "data": "Review added successfully!"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
