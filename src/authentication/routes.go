package authentication

import (
	"github.com/gin-gonic/gin"
)

func Routes() {
	server := gin.Default()

	server.GET("/token", CheckToken)
}
