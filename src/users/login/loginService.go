package users

import (
	"context"
	"fmt"

	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var myCollection *mongo.Collection = database.OpenCollection(database.Client, "users")

func LoginService(login *entity.Credentials) string {
	// validate credentials
	user := entity.User{}
	query := bson.M{"email": login.Email}
	err2 := myCollection.FindOne(context.TODO(), query).Decode(&user)
	if err2 != nil {
		return "Error"
	}

	// Comparing the password with the hash
	loginPass := []byte(login.Password)
	queryPass := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(queryPass, loginPass)
	fmt.Println(err) // nil means it is a match

	if err != nil {
		return "Error"
	} else {
		signToken := auth.SignToken(login.Email, 60)
		return signToken
	}
}
