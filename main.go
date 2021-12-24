package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	contents "gitlab.com/JacobDCruz/supplier-portal/src/contents/controllers"
	delete "gitlab.com/JacobDCruz/supplier-portal/src/person/delete"
	get "gitlab.com/JacobDCruz/supplier-portal/src/person/get"
	list "gitlab.com/JacobDCruz/supplier-portal/src/person/list"

	// profile
	getProfile "gitlab.com/JacobDCruz/supplier-portal/src/profiles/get"
	updateProfile "gitlab.com/JacobDCruz/supplier-portal/src/profiles/update"

	// users
	listUser "gitlab.com/JacobDCruz/supplier-portal/src/users/list"
	loginUser "gitlab.com/JacobDCruz/supplier-portal/src/users/login"
	logoutUser "gitlab.com/JacobDCruz/supplier-portal/src/users/logout"
	signupUser "gitlab.com/JacobDCruz/supplier-portal/src/users/signup"

	// products
	addProduct "gitlab.com/JacobDCruz/supplier-portal/src/products/add"
	getProduct "gitlab.com/JacobDCruz/supplier-portal/src/products/get"
	listProduct "gitlab.com/JacobDCruz/supplier-portal/src/products/list"

	// votes
	votes "gitlab.com/JacobDCruz/supplier-portal/src/votes/controllers"
)

func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

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
	server.POST("/google/login", loginUser.GoogleLogin)

	// users
	server.POST("/login", func(ctx *gin.Context) {
		loginUser.LoginController(ctx)
	})
	server.POST("/logout", logoutUser.LogoutController)
	server.GET("/users", func(ctx *gin.Context) {
		listUser.ListController(ctx)
	})
	server.POST("/signup", func(ctx *gin.Context) {
		signupUser.SignupController(ctx)
	})

	// profile
	server.PUT("/profile/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		updateProfile.UpdateController(ctx, id)
	})

	// contents
	server.GET("/content", contents.ListController)
	server.GET("/content/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		getProfile.GetController(ctx, id)
	})
	server.POST("/content", contents.AddController)

	// products
	server.GET("/products", listProduct.ListController)
	server.GET("/product", func(ctx *gin.Context) {
		getProduct.GetController(ctx)
	})
	server.POST("/products", addProduct.AddController)

	// votations
	server.GET("/vote/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		votes.GetController(ctx, id)
	})
	server.POST("/vote", votes.AddController)

	// register
	// server.Use(CORSMiddleware())
	server.Run(":8010")
}
