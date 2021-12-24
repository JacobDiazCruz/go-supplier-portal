package products

import (
	"context"
	"encoding/json"
	"fmt"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/products/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

type getService interface {
	GetService() entity.Product
}

type Param struct {
	id   string
	slug string
}

func (p Param) GetService() entity.Product {
	// set initial values
	result := entity.Product{}
	var query = bson.M{"_id": ""}
	objID, err := primitive.ObjectIDFromHex(p.id)

	// _id query params
	if err != nil {
		fmt.Println("no id found")
	} else {
		query = bson.M{"_id": objID}
		fmt.Println(query)
	}

	// slug query params
	if p.slug != "" {
		query = bson.M{"slug": p.slug}
	}

	// query to db
	err2 := productCollection.FindOne(context.TODO(), query).Decode(&result)
	if err2 != nil {
		panic(err2)
	}

	// log documents and return
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	fmt.Println("Test123123")
	fmt.Println(result)
	return result
}
