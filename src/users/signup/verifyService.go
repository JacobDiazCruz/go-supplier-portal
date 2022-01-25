package users

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"strings"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var codeCollection *mongo.Collection = database.OpenCollection(database.Client, "signup_codes")

type VerifyCode struct {
	Code string `json:"code" bson:"code"`
}

func MakeVerificationCode() string {
	// generate code
	b := make([]byte, 2) //equals 6 characters
	rand.Read(b)
	s := hex.EncodeToString(b)
	code := strings.ToUpper(s)

	// save generated code to db
	result, err := codeCollection.InsertOne(context.TODO(), bson.M{
		"code": code,
	})
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}

func GetVerificationCode(code string) VerifyCode {
	codeStruct := VerifyCode{}

	// query
	query := bson.M{"email": code}
	codeCollection.FindOne(context.TODO(), query).Decode(&codeStruct)
	return codeStruct
}
