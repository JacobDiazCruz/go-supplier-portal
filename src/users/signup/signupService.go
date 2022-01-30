package users

import (
	"context"
	"fmt"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func SignupService(user entity.User, signupType string) string {
	// initialize not verified user
	if signupType == "manual" {
		user.Verified = false
	} else {
		user.Verified = true
	}

	fmt.Println(user)
	fmt.Println("hehehee")

	// save user to db
	result, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
