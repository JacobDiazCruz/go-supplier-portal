package votes

import (
	"context"
	"log"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/votes/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var voteCollection *mongo.Collection = database.OpenCollection(database.Client, "votes")

func ListService() []entity.Vote {
	// query list
	cursor, err := voteCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	votes := []entity.Vote{}
	if err = cursor.All(context.TODO(), &votes); err != nil {
		log.Fatal(err)
	}
	return votes
}
