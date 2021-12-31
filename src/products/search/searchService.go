package products

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

func SearchService(keyword string) []bson.M {
	// filter := bson.M{
	// 	"$search": bson.D{{"query": keyword}},
	// }
	var result []bson.M
	fmt.Println(keyword)
	fmt.Println("yeyeyeyeyeyeyyee")
	query := bson.M{
		"$and": []bson.M{
			bson.M{"$or": []bson.M{
				bson.M{"slug": bson.M{"$regex": keyword, "$options": "i"}},
			}},
		},
	}

	cursor, err := productCollection.Find(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonData)

	return result
}
