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

func CancelService(orderId string) string {
	// cancel status
	objID, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		panic(err)
	}

	orderStatus := entity.OrderStatus{}
	orderStatus.Title = "Order Cancelled"
	orderStatus.Label = "cancelled"

	// query filters
	filter := bson.M{"_id": objID}
	update := bson.M{"order_status.title": orderStatus.Title, "order_status.label": orderStatus.Label}

	// query db
	res := orderCollection.FindOneAndUpdate(context.Background(),
		filter,
		bson.M{"$set": update},
	)

	// check error
	if res.Err() != nil {
		panic(res.Err())
	}

	// return if no error
	return "Success"
}
