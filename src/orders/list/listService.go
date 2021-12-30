package orders

import (
	"context"
	"log"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/orders/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var addressCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

type listService interface {
	ListService() entity.Order
}

func ListService(userId primitive.ObjectID) []entity.Order {
	cursor, err := addressCollection.Find(context.TODO(), bson.M{
		"cart.user_id": userId,
	})
	if err != nil {
		log.Fatal(err)
	}
	addresses := []entity.Order{}
	if err = cursor.All(context.TODO(), &addresses); err != nil {
		log.Fatal(err)
	}
	return addresses
}
