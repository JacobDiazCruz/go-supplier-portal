package profiles

import (
	"context"
	"log"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/profiles/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var profileCollection *mongo.Collection = database.OpenCollection(database.Client, "profiles")

func ListService() []entity.Profile {
	cursor, err := profileCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	profiles := []entity.Profile{}
	if err = cursor.All(context.TODO(), &profiles); err != nil {
		log.Fatal(err)
	}
	return profiles
}
