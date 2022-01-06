package addresses

import (
	"context"
	"encoding/json"
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
	addresses := []entity.Address{}
	var result []bson.M
	cursor, err := addressCollection.Find(context.TODO(), bson.M{
		"user_id": userId,
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
	json.Unmarshal(jsonData, &addresses)
	return addresses
}
