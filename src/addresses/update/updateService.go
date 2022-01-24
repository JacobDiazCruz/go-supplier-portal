package addresses

import (
	"context"
	"fmt"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/addresses/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var addressCollection *mongo.Collection = database.OpenCollection(database.Client, "addresses")

func UpdateService(address entity.Address, id string) string {
	// convert id string to mongo
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	// query
	result, err := addressCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		address,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	return "Success"
}
