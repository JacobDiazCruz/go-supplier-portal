package profiles

import (
	"context"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/profiles/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var profileCollection *mongo.Collection = database.OpenCollection(database.Client, "profiles")

func AddService(profile entity.Profile) string {
	// query
	result, err := profileCollection.InsertOne(context.TODO(), bson.M{
		"email":     profile.Email,
		"firstname": profile.FirstName,
		"lastname":  profile.LastName,
		"role":      profile.Role,
		"user_id":   profile.UserId,
	})
	if err != nil {
		panic(err)
	}

	// return string id
	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
