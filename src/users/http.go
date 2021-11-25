package users

import (
	"github.com/gin-gonic/gin"
	get "gitlab.com/JacobDCruz/supplier-portal/src/person/get"
	users "gitlab.com/JacobDCruz/supplier-portal/src/users/controllers"
)

func Routes() {
	server := gin.Default()

	server.POST("/login", users.LoginController)
	server.POST("/refresh", users.RefreshController)

	server.GET("/users", users.ListController)

	server.POST("/signup", func(ctx *gin.Context) {
		id := users.SignupController(ctx)
		get.GetUser(ctx, id)
	})

	// register
	server.Run(":8000")
}
