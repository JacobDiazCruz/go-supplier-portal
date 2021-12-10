package profiles

import (
	"context"
	"fmt"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/profiles/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddService(profile entity.Profile) string {
	fmt.Println(profile)
	fmt.Println("add service here id ^")
	// query
	result, err := profileCollection.InsertOne(context.TODO(), profile)
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
