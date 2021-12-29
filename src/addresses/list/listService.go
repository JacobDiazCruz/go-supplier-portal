package addresses

import (
	"context"
	"log"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/addresses/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var addressCollection *mongo.Collection = database.OpenCollection(database.Client, "addresses")

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
