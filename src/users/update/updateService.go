package users

import (
	"context"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func UpdateService(user entity.User) string {
	// query
	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	}

	// query db
	res := userCollection.FindOneAndUpdate(context.Background(),
		filter,
		bson.M{"$set": update},
	)

	// check error
	if res.Err() != nil {
		panic(res.Err())
	}

	// return if no error
	return "Success"
}
