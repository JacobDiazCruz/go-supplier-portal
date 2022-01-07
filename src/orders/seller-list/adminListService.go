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
)

var addressCollection *mongo.Collection = database.OpenCollection(database.Client, "seller_orders")

type listService interface {
	ListService() entity.Order
}

func ListService(userId primitive.ObjectID) []entity.Order {
	orders := []entity.Order{}
	var result []bson.M
	cursor, err := addressCollection.Find(context.TODO(), bson.M{
		"cart.user_id": userId,
	})
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
