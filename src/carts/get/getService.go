package carts

import (
	"context"
	"encoding/json"
	"fmt"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "carts")
var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

type Params struct {
	ProductIds []primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
}

func GetService(userId primitive.ObjectID) entity.GetCart {
	result := entity.GetCart{}

	// query
	query := bson.M{"user_id": userId}
	err := cartCollection.FindOne(context.TODO(), query).Decode(&result)
	jsonData, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	fmt.Println("Test123123")
	fmt.Println(result)
	return result
}
