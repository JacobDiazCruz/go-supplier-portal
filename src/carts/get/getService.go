package carts

import (
	"context"
	"encoding/json"

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
	cart := entity.GetCart{}
	var result bson.M

	// query
	query := bson.M{"user_id": userId}
	err2 := cartCollection.FindOne(context.TODO(), query).Decode(&result)
	if err2 != nil {
		panic(err2)
	}

	// unmarshal result to products struct
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &cart)

	return cart
}
