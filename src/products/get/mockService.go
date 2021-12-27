package products

import (
	"context"
	"encoding/json"
	"fmt"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/products/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockGetService interface {
	GetService() entity.Product
}

type MockParam struct {
	id   string
	slug string
}

func (p MockParam) MockGetService() entity.Product {
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
