package carts

import (
	"context"
	"fmt"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "carts")

func ClearService(id string) string {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	// query
	query := bson.M{"_id": objID}
	result, err := cartCollection.DeleteOne(context.TODO(), query)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	return "Success"
}
