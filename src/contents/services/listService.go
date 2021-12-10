package contents

import (
	"context"
	"log"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/contents/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var contentCollection *mongo.Collection = database.OpenCollection(database.Client, "contents")

func ListService() []entity.Content {
	cursor, err := contentCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	contents := []entity.Content{}
	if err = cursor.All(context.TODO(), &contents); err != nil {
		log.Fatal(err)
	}
	return contents
}
