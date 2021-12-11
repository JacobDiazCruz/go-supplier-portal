package profiles

import (
	"context"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/profiles/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
