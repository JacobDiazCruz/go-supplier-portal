package users

import (
	"context"
	"log"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func ListService() []entity.User {
	cursor, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	users := []entity.User{}
	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Fatal(err)
	}
	return users
}
