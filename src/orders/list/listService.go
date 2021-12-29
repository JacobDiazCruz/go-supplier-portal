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
	ListService() entity.Address
}

func ListService(userId primitive.ObjectID) []entity.Address {
	cursor, err := addressCollection.Find(context.TODO(), bson.M{
		"user_id": userId,
	})
	if err != nil {
		log.Fatal(err)
	}
	addresses := []entity.Address{}
	if err = cursor.All(context.TODO(), &addresses); err != nil {
		log.Fatal(err)
	}
	return addresses
}
