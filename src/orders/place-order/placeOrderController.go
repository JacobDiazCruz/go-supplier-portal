package orders

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/orders/entity"
	user "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
)

func PlaceOrderController(ctx *gin.Context) {
	// check token and return
	ct := auth.GetToken(ctx)

	b := make([]byte, 7) //equals 14 characters
	rand.Read(b)
	s := hex.EncodeToString(b)
	randOrderId := strings.ToUpper(s)

	// if no error
	if ct != nil {
		// get email and return user details
		u := user.GetEmail(ct.Email)
		au := entity.Auth{}
		au.UserId = u.ID

		// cart request
		order := entity.PlaceOrder{}
		err := ctx.BindJSON(&order)
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
		}
		order.OrderId = randOrderId
		order.UserId = u.ID

		// audit log
		auditLog := &order.AuditLog
		auditLog.Name = ct.Username
		auditLog.Email = ct.Email
		auditLog.ThumbnailImage = ct.ThumbnailImage
		auditLog.OriginalImage = ct.OriginalImage
		auditLog.CreatedAt = time.Now()
		auditLog.CreatedBy = ct.Username
		auditLog.UpdatedAt = time.Now()
		auditLog.UpdatedBy = ct.Username

		// update service
		res := PlaceOrderService(order, au)
		fmt.Println(res)

		ctx.JSON(http.StatusOK, gin.H{"msg": "Placed Order Successfully!", "data": res})
	} else { // if error exist
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Token"})
	}
}
