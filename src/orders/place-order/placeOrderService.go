package orders

import (
	"context"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/orders/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

// create order id and return
func PlaceOrderService(order entity.PlaceOrder) string {
	// cartid string to mongoid
	cartId, err := primitive.ObjectIDFromHex(order.CartId)
	if err != nil {
		panic(err)
	}

	// query
	result, err := orderCollection.InsertOne(context.TODO(), bson.M{
		"cart_id":          cartId,
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
