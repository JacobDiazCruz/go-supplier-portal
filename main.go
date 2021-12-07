package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	delete "gitlab.com/JacobDCruz/supplier-portal/src/person/delete"
	get "gitlab.com/JacobDCruz/supplier-portal/src/person/get"
	list "gitlab.com/JacobDCruz/supplier-portal/src/person/list"
	users "gitlab.com/JacobDCruz/supplier-portal/src/users/controllers"
)

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		ctx.Header("Access-Control-Allow-Origin", "*")
// 		ctx.Header("Access-Control-Allow-Credentials", "true")
// 		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type, Authorization")
// 		ctx.Header("X-Frame-Options", "sameorigin")
// 		ctx.Header("Content-Security-Policy", "self")
// 		ctx.Header("X-Content-Type-Options", "nosniff")
// 		ctx.Header("Content-Type", "multipart/form-data")
// 		ctx.Header("Referrer-Policy", "origin")
// 		ctx.Header("X-XSS-Protection", "1; mode=block")
// 		ctx.Header("Accept", "image/*")

// 		if ctx.Request.Method == "OPTIONS" {
// 			return
// 		}

// 		ctx.Next()
// 	}
// }

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
	server.POST("/google/login", auth.GoogleLogin)
	// server.POST("/google/callback", auth.GoogleCallback)

	// users
	server.POST("/login", func(ctx *gin.Context) {
		users.LoginController(ctx)
	})
	server.POST("/refresh", users.RefreshController)
	server.POST("/logout", users.LogoutController)
	server.GET("/users", func(ctx *gin.Context) {
		users.ListController(ctx)
	})

	server.POST("/signup", func(ctx *gin.Context) {
		id := users.SignupController(ctx)
		get.GetUser(ctx, id)
	})

	// register
	// server.Use(CORSMiddleware())
	server.Run(":8000")
}
