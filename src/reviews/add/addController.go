package reviews

import (
	"fmt"
	"net/http"
	"time"

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

		// audit log
		auditLog := &review.AuditLog
		auditLog.Name = ct.Username
		auditLog.Email = ct.Email
		auditLog.ThumbnailImage = ct.ThumbnailImage
		auditLog.OriginalImage = ct.OriginalImage
		auditLog.CreatedAt = time.Now()
		auditLog.CreatedBy = ct.Username
		auditLog.UpdatedAt = time.Now()
		auditLog.UpdatedBy = ct.Username

		res := AddService(review)
		fmt.Println(res)
		// return res
		ctx.JSON(http.StatusOK, gin.H{"msg": "Product added successfully.", "data": "Review added successfully!"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
