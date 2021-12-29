package orders

import (
	"context"
	"fmt"

	getCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/get"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/orders/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

// create order id and return
func PlaceOrderService(order entity.PlaceOrder, au entity.Auth) string {
	// cartid string to mongoid
	cartId, err := primitive.ObjectIDFromHex(order.CartId)
	if err != nil {
		panic(err)
	}

	cartRes := getCart.GetService(au.UserId)
	// vals := make([]string, 10)
	for _, val := range cartRes.Products {
		for _, val2 := range val {
			// md := val2.(map[string]interface{})
			fmt.Println("shit")
			fmt.Println(val2)
			fmt.Println("shit")
			// vals = append(vals, val2)
		}
	}
	// fmt.Println(vals)
	fmt.Println("im here wooo1111")
	fmt.Println(cartRes.Products)
	fmt.Println("im here wooo222")

	// query
	result, err := orderCollection.InsertOne(context.TODO(), bson.M{
		"cart_id":          cartId,
		"user_id":          order.UserId,
		"delivery_address": order.DeliveryAddress,
		"audit_log":        order.AuditLog,
	})
	if err != nil {
		panic(err)
	}

	// return order id string
	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
