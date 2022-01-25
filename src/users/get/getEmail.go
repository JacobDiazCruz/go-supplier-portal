package users

import (
	"context"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func GetEmail(email string) entity.User {
	result := entity.User{}

	// query
	query := bson.M{"email": email, "verified": true}
	userCollection.FindOne(context.TODO(), query).Decode(&result)
	return result
}
