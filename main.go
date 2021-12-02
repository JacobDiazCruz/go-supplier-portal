package main

import (
	"github.com/gin-gonic/gin"
	delete "gitlab.com/JacobDCruz/supplier-portal/src/person/delete"
	get "gitlab.com/JacobDCruz/supplier-portal/src/person/get"
	list "gitlab.com/JacobDCruz/supplier-portal/src/person/list"
	users "gitlab.com/JacobDCruz/supplier-portal/src/users/controllers"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/authentication"
)

func main() {
	server := gin.Default()

	server.GET("/person/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		get.GetUser(ctx, id)
	})

	server.GET("/person", func(ctx *gin.Context) {
		ctx.JSON(200, list.GetAllUsers())
	})

	server.DELETE("/person/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		delete.DeleteController(ctx, id)
	})

	// auth
	server.GET("/token", auth.CheckToken)

	// users
	v1 := server.Group("/")
	{
		v1.POST("/login", users.LoginController)
		v1.POST("/refresh", users.RefreshController)
		v1.POST("/logout", users.LogoutController)
		v1.GET("/users", users.ListController)

		v1.POST("/signup", func(ctx *gin.Context) {
			id := users.SignupController(ctx)
			get.GetUser(ctx, id)
		})
	}

	// register
	server.Run(":8000")
}
