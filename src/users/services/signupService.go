package users

import (
	"context"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func SignupService(user entity.User) string {
	// query
	result, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	// test := map[string]string{
	// 	"id": oid.Hex(),
	// }
	// return get.GetUser(oid.Hex())
	return oid.Hex()
}
