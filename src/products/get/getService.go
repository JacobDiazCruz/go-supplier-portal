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

func GetService(id string, slug string) entity.Product {
	// set initial values
	product := entity.Product{}
	var result bson.M

	var query = bson.M{"_id": ""}
	objID, err := primitive.ObjectIDFromHex(id)

	// _id query params
	if err != nil {
		fmt.Println("no id found")
	} else {
		query = bson.M{"_id": objID}
		fmt.Println(query)
	}

	// slug query params
	if slug != "" {
		query = bson.M{"slug": slug}
	}

	// query to db
	err2 := productCollection.FindOne(context.TODO(), query).Decode(&result)
	if err2 != nil {
		panic(err2)
	}

	// unmarshal result to products struct
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &product)
	return product
}

func GetVariantOptions(id string) entity.Product {
	// set initial values
	product := entity.Product{}
	var result bson.M

	var query = bson.M{"_id": ""}
	objID, err := primitive.ObjectIDFromHex(id)

	// _id query params
	if err != nil {
		fmt.Println("no id found")
	} else {
		query = bson.M{"_id": objID}
		fmt.Println(query)
	}

	// query to db
	err2 := productCollection.FindOne(context.TODO(), query).Decode(&result)
	if err2 != nil {
		panic(err2)
	}

	// unmarshal result to products struct
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &product)
	return product
}
