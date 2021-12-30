package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	delete "gitlab.com/JacobDCruz/supplier-portal/src/person/delete"
	get "gitlab.com/JacobDCruz/supplier-portal/src/person/get"
	list "gitlab.com/JacobDCruz/supplier-portal/src/person/list"

	// profile
	updateProfile "gitlab.com/JacobDCruz/supplier-portal/src/profiles/update"

	// addresses
	addAddress "gitlab.com/JacobDCruz/supplier-portal/src/addresses/add"
	getAddress "gitlab.com/JacobDCruz/supplier-portal/src/addresses/get"
	listAddress "gitlab.com/JacobDCruz/supplier-portal/src/addresses/list"
	updateAddress "gitlab.com/JacobDCruz/supplier-portal/src/addresses/update"

	// users
	listUser "gitlab.com/JacobDCruz/supplier-portal/src/users/list"
	loginUser "gitlab.com/JacobDCruz/supplier-portal/src/users/login"
	logoutUser "gitlab.com/JacobDCruz/supplier-portal/src/users/logout"
	signupUser "gitlab.com/JacobDCruz/supplier-portal/src/users/signup"

	// products
	addProduct "gitlab.com/JacobDCruz/supplier-portal/src/products/add"
	adminListProduct "gitlab.com/JacobDCruz/supplier-portal/src/products/admin-list"
	getProduct "gitlab.com/JacobDCruz/supplier-portal/src/products/get"
	listProduct "gitlab.com/JacobDCruz/supplier-portal/src/products/list"

	// carts
	addCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/add"
	// clearCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/clear-cart"
	deleteCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/delete-item"
	getCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/get"
	updateCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/update"

	// orders
	listOrder "gitlab.com/JacobDCruz/supplier-portal/src/orders/list"
	getOrder "gitlab.com/JacobDCruz/supplier-portal/src/orders/get"
	placeOrder "gitlab.com/JacobDCruz/supplier-portal/src/orders/place-order"
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

	// products
	server.GET("/products", listProduct.ListController)
	server.GET("/admin-products", adminListProduct.AdminListController)
	server.GET("/product", func(ctx *gin.Context) {
		getProduct.GetController(ctx)
	})
	server.POST("/products", addProduct.AddController)

	// carts
	server.GET("/cart/get", func(ctx *gin.Context) {
		getCart.GetController(ctx)
	})
	server.POST("/cart/add", func(ctx *gin.Context) {
		addCart.AddController(ctx)
	})
	server.PUT("/remove-item", func(ctx *gin.Context) {
		deleteCart.DeleteController(ctx)
	})
	server.PUT("/update-quantity", func(ctx *gin.Context) {
		updateCart.UpdateController(ctx)
	})

	// addresses
	server.GET("/address/get/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		getAddress.GetController(ctx, id)
	})
	server.GET("/address/list", func(ctx *gin.Context) {
		listAddress.ListController(ctx)
	})
	server.POST("/address/add", func(ctx *gin.Context) {
		addAddress.AddController(ctx)
	})
	server.PUT("/address/update/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		updateAddress.UpdateController(ctx, id)
	})

	// orders
	server.GET("/orders/list", func(ctx *gin.Context) {
		listOrder.ListController(ctx)
	})
	server.GET("/orders/get/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		getOrder.GetController(ctx, id)
	})
	server.POST("/orders/place", func(ctx *gin.Context) {
		placeOrder.PlaceOrderController(ctx)
	})

	// register
	// server.Use(CORSMiddleware())
	server.Run(":8010")
}
