package users

import (
	"context"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func GetEmail(email string) (entity.User, error) {
	result := entity.User{}

	// query
	query := bson.M{"email": email}
	userCollection.FindOne(context.TODO(), query).Decode(&result)
	return result, nil
}
