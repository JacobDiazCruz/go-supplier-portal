package products

import (
	"context"
	"encoding/json"
	"log"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/products/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

func AdminListService(listFilters entity.List) []entity.Product {
	products := []entity.Product{}
	var result []bson.M

	options := options.Find()
	options.SetLimit(listFilters.Limit)
	cursor, err := productCollection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}

	// unmarshal result to products struct
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &products)
	return products
}
