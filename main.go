// package main

// import (
// 	"log"
// 	"net/http"

// 	auth "gitlab.com/JacobDCruz/supplier-portal/src/authentication"
// )

// func main() {
// 	http.HandleFunc("/login", auth.Login)
// 	http.HandleFunc("/home", auth.Home)
// 	http.HandleFunc("/refresh", auth.Refresh)

// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

package main

import (
	"github.com/gin-gonic/gin"
	delete "gitlab.com/JacobDCruz/supplier-portal/src/person/delete"
	get "gitlab.com/JacobDCruz/supplier-portal/src/person/get"
	list "gitlab.com/JacobDCruz/supplier-portal/src/person/list"
	users "gitlab.com/JacobDCruz/supplier-portal/src/users"
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

	users.Routes()

	server.DELETE("/person/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		delete.DeleteController(ctx, id)
	})

	// register
	server.Run(":8000")
}
