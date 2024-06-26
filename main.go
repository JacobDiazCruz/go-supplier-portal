package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"

	// profile
	updateProfile "gitlab.com/JacobDCruz/supplier-portal/src/profiles/update"

	// addresses
	addAddress "gitlab.com/JacobDCruz/supplier-portal/src/addresses/add"
	getAddress "gitlab.com/JacobDCruz/supplier-portal/src/addresses/get"
	listAddress "gitlab.com/JacobDCruz/supplier-portal/src/addresses/list"
	updateAddress "gitlab.com/JacobDCruz/supplier-portal/src/addresses/update"

	// users
	changePassword "gitlab.com/JacobDCruz/supplier-portal/src/users/change-password"
	forgotPassword "gitlab.com/JacobDCruz/supplier-portal/src/users/forgot-password"
	getUser "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
	listUser "gitlab.com/JacobDCruz/supplier-portal/src/users/list"
	loginUser "gitlab.com/JacobDCruz/supplier-portal/src/users/login"
	logoutUser "gitlab.com/JacobDCruz/supplier-portal/src/users/logout"
	resetPassword "gitlab.com/JacobDCruz/supplier-portal/src/users/reset-password"
	signupUser "gitlab.com/JacobDCruz/supplier-portal/src/users/signup"
	updateUser "gitlab.com/JacobDCruz/supplier-portal/src/users/update"

	// products
	addProduct "gitlab.com/JacobDCruz/supplier-portal/src/products/add"
	adminListProduct "gitlab.com/JacobDCruz/supplier-portal/src/products/admin-list"
	getProduct "gitlab.com/JacobDCruz/supplier-portal/src/products/get"
	listProduct "gitlab.com/JacobDCruz/supplier-portal/src/products/list"
	searchProduct "gitlab.com/JacobDCruz/supplier-portal/src/products/search"

	// reviews
	addReview "gitlab.com/JacobDCruz/supplier-portal/src/reviews/add"
	listReview "gitlab.com/JacobDCruz/supplier-portal/src/reviews/list"

	// carts
	addCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/add"
	clearCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/clear-cart"
	deleteCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/delete-item"
	getCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/get"
	updateCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/update"
	validateCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/validate-items"

	// orders
	addSellerOrder "gitlab.com/JacobDCruz/supplier-portal/src/orders/add-seller-order"
	cancelOrder "gitlab.com/JacobDCruz/supplier-portal/src/orders/cancel"
	getOrder "gitlab.com/JacobDCruz/supplier-portal/src/orders/get"
	listOrder "gitlab.com/JacobDCruz/supplier-portal/src/orders/list"
	ordersPaymongo "gitlab.com/JacobDCruz/supplier-portal/src/orders/paymongo"
	placeOrder "gitlab.com/JacobDCruz/supplier-portal/src/orders/place-order"
)

func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	// auth
	server.POST("/token", auth.CheckToken)
	server.POST("/google/login", loginUser.GoogleLogin)
	server.POST("/facebook/login", loginUser.FacebookLogin)

	// users
	server.POST("/login", func(ctx *gin.Context) {
		loginUser.LoginController(ctx)
	})
	server.POST("/reset-password", func(ctx *gin.Context) {
		resetPassword.ResetController(ctx)
	})
	server.POST("/user/change-password", func(ctx *gin.Context) {
		changePassword.ChangeController(ctx)
	})
	server.POST("/logout", logoutUser.LogoutController)
	server.GET("/users", func(ctx *gin.Context) {
		listUser.ListController(ctx)
	})
	server.GET("/user/get/:email", func(ctx *gin.Context) {
		email := ctx.Param("email")
		getUser.GetController(ctx, email)
	})
	server.PUT("/user/update", func(ctx *gin.Context) {
		updateUser.UpdateController(ctx)
	})
	server.POST("/signup", func(ctx *gin.Context) {
		signupUser.SignupController(ctx)
	})
	server.POST("/verify-email", func(ctx *gin.Context) {
		signupUser.VerifyEmailController(ctx)
	})
	server.POST("/verify-code", func(ctx *gin.Context) {
		signupUser.VerifyCodeController(ctx)
	})
	server.POST("/forgot-password", forgotPassword.ForgotController)

	// profile
	server.PUT("/profile/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		updateProfile.UpdateController(ctx, id)
	})

	// paymongo
	server.POST("/checkout", ordersPaymongo.PaymongoService)

	// products
	server.GET("/products", listProduct.ListController)
	server.GET("/admin-products", adminListProduct.ListController)
	server.GET("/product", func(ctx *gin.Context) {
		getProduct.GetController(ctx)
	})
	server.POST("/products", addProduct.AddController)
	server.POST("/products/search", searchProduct.SearchController)

	// reviews
	server.POST("/reviews/add", addReview.AddController)
	server.GET("/reviews/list", listReview.ListController)

	// carts
	server.GET("/cart/get", func(ctx *gin.Context) {
		getCart.GetController(ctx)
	})
	server.POST("/cart/add", func(ctx *gin.Context) {
		addCart.AddController(ctx)
	})
	server.POST("/cart/clear", func(ctx *gin.Context) {
		clearCart.ClearController(ctx)
	})
	server.PUT("/remove-item", func(ctx *gin.Context) {
		deleteCart.DeleteController(ctx)
	})
	server.PUT("/update-quantity", func(ctx *gin.Context) {
		updateCart.UpdateController(ctx)
	})
	server.GET("/cart/validate-items", func(ctx *gin.Context) {
		validateCart.ValidateController(ctx)
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
	server.POST("/orders/add-seller-order", func(ctx *gin.Context) {
		addSellerOrder.AddController(ctx)
	})
	server.PUT("/orders/cancel/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		cancelOrder.CancelController(ctx, id)
	})

	// register
	// server.Use(CORSMiddleware())
	server.Run(":8010")
}
