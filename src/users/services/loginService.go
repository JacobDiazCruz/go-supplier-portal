package users

import (
	"context"

	auth "gitlab.com/JacobDCruz/supplier-portal/src/authentication"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var myCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func LoginService(login *entity.Credentials) string {
	result := entity.User{}

	// validate credentials
	query := bson.M{"email": login.Email, "password": login.Password}
	err := myCollection.FindOne(context.TODO(), query).Decode(&result)
	if err != nil {
		return "Incorrent email or password"
	}

	// sign token and return
	tk := auth.TokenIdentity{}
	signToken := auth.SignToken(tk)
	return signToken
}
