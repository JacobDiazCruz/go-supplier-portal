package profiles

import (
	"context"
	"fmt"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/profiles/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var profileCollection *mongo.Collection = database.OpenCollection(database.Client, "profiles")

func UpdateService(profile entity.Profile, profileId string) string {
	// convert id string to mongo
	objID, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		panic(err)
	}

	// query
	result, err := profileCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"email":     profile.Email,
				"firstname": profile.FirstName,
				"lastname":  profile.LastName,
				"role":      profile.Role,
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return "Success"
}
