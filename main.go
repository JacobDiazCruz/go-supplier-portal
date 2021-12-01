package main

import (
	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/authentication"
	delete "gitlab.com/JacobDCruz/supplier-portal/src/person/delete"
	get "gitlab.com/JacobDCruz/supplier-portal/src/person/get"
	list "gitlab.com/JacobDCruz/supplier-portal/src/person/list"
	users "gitlab.com/JacobDCruz/supplier-portal/src/users"
)

func main() {
	server := gin.Default()

	server.GET("/token", auth.CheckToken)
	server.GET("/person/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		get.GetUser(ctx, id)
	})

	server.GET("/person", func(ctx *gin.Context) {
		ctx.JSON(200, list.GetAllUsers())
	})

	users.Routes()

	server.DELETE("/person/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		delete.DeleteController(ctx, id)
	})
}
