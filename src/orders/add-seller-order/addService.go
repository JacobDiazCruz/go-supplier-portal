package orders

import (
	"context"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/orders/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "seller_orders")

func AddService(order entity.SellerOrder) string {
	// query
	result, err := orderCollection.InsertOne(context.TODO(), bson.M{
		"product":      order.Product,
		"order_id":     order.OrderId,
		"order_status": order.OrderStatus,
		"quantity":     order.Quantity,
		"audit_log":    order.AuditLog,
	})
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
