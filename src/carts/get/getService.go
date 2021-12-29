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
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	fmt.Println("Test123123")
	fmt.Println(result)
	return result
}

// func ListService() []entity.Product {
// 	// Get Cart list
// 	cursor, err := cartCollection.FindOne(context.TODO(), query).bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	carts := []entity.Cart{}
// 	if err = cursor.All(context.TODO(), &carts); err != nil {
// 		log.Fatal(err)
// 	}
// 	products := entity.Cart{}

// 	// append to productids
// 	var param = Params{}
// 	for i, s := range carts.Products {
// 		fmt.Println(i, s.ProductId)
// 		param.ProductIds = append(param.ProductIds, s.ProductId)
// 	}

// 	// Get Product list
// 	cursor2, err2 := productCollection.Find(context.TODO(), bson.M{"_id": bson.M{"$in": param.ProductIds}})
// 	if err2 != nil {
// 		log.Fatal(err2)
// 	}
// 	products := []entity.Product{}
// 	if err = cursor2.All(context.TODO(), &products); err != nil {
// 		log.Fatal(err)
// 	}
// 	return products
// }
