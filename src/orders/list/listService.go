package orders

import (
	"context"
	"encoding/json"
	"log"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/orders/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var addressCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

type listService interface {
	ListService() entity.Order
}

func ListService(userId primitive.ObjectID) []entity.Order {
	orders := []entity.Order{}
	var result []bson.M

	// Sort filter
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"audit_log.created_at", -1}})

	// Find Query
	cursor, err := addressCollection.Find(context.TODO(), bson.M{
		"cart.user_id": userId,
	}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &orders)
	return orders
}
