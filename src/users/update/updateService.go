package users

import (
	"context"
	"fmt"

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
		"email":          user.Email,
		"username":       user.Username,
		"contact_number": user.ContactNumber,
	}

	// query db
	res, err := userCollection.UpdateOne(context.TODO(),
		filter,
		bson.M{"$set": update},
	)

	// check error
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	// return if no error
	return "Success"
}
