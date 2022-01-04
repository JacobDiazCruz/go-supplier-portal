package reviews

import (
	"context"
	"encoding/json"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/reviews/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var reviewCollection *mongo.Collection = database.OpenCollection(database.Client, "reviews")

type Param struct {
	id string
}

func GetService(id string) entity.Review {
	review := entity.Review{}
	var result bson.M
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	// query
	query := bson.M{"_id": objID}
	if err2 := reviewCollection.FindOne(context.TODO(), query).Decode(&result); err != nil {
		panic(err2)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &review)
	return review
}
