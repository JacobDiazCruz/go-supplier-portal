package products

import (
	"context"
	"log"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/products/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

func ListService() []entity.Product {
	cursor, err := productCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	products := []entity.Product{}
	if err = cursor.All(context.TODO(), &products); err != nil {
		log.Fatal(err)
	}
	return products
}
