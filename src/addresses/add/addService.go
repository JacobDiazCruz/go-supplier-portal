package addresses

import (
	"context"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/addresses/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var addressesCollection *mongo.Collection = database.OpenCollection(database.Client, "addresses")

func AddService(address entity.Address) string {
	// query
	result, err := addressesCollection.InsertOne(context.TODO(), address)
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
