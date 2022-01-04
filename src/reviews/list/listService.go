package reviews

import (
	"context"
	"log"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/reviews/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var reviewCollection *mongo.Collection = database.OpenCollection(database.Client, "reviews")

func ListService(listFilters entity.Review, productId string) []entity.Review {
	// mongo id
	objID, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		panic(err)
	}

	// query db
	cursor, err := reviewCollection.Find(context.TODO(), bson.M{
		"product_id": objID,
	})
	if err != nil {
		log.Fatal(err)
	}
	reviews := []entity.Review{}
	if err = cursor.All(context.TODO(), &reviews); err != nil {
		log.Fatal(err)
	}
	return reviews
}
