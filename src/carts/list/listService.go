package carts

import (
	"context"
	"fmt"
	"log"

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

func ListService() []entity.Product {
	// Get Cart list
	cursor, err := cartCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	carts := []entity.Cart{}
	if err = cursor.All(context.TODO(), &carts); err != nil {
		log.Fatal(err)
	}

	// append to productids
	var param = Params{}
	for i, s := range carts {
		fmt.Println(i, s.ProductId)
		param.ProductIds = append(param.ProductIds, s.ProductId)
	}

	// Get Product list
	// ERROR HERE
	fmt.Println(param.ProductIds)
	fmt.Println("gwegwegwegwggwe")
	cursor2, err2 := productCollection.Find(context.TODO(), bson.M{"_id": bson.M{"$in": param.ProductIds}})
	if err2 != nil {
		log.Fatal(err2)
	}
	products := []entity.Product{}
	if err = cursor2.All(context.TODO(), &products); err != nil {
		log.Fatal(err)
	}
	return products
}
