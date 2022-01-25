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
	// save user to db
	result, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
