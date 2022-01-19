package products

import (
	"context"
	"encoding/json"
	"log"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/products/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

func SearchService(sf SearchField) []bson.M {
	// init
	products := []entity.Product{}
	var result []bson.M

	// query filters
	query := bson.M{
		"$and": []bson.M{
			bson.M{"$or": []bson.M{
				bson.M{"slug": bson.M{"$regex": sf.Search, "$options": "i"}},
			}},
		},
	}

	// query db
	cursor, err := productCollection.Find(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}

	// convert to json struct and return
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &products)
	return result
}
