package users

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var codeCollection *mongo.Collection = database.OpenCollection(database.Client, "signup_codes")

type Verification struct {
	Email string `json:"email" bson:"email"`
	Code  string `json:"code" bson:"code"`
}

func MakeVerificationCode(email string) string {
	// generate code
	b := make([]byte, 3) //equals 6 characters
	rand.Read(b)
	s := hex.EncodeToString(b)
	code := strings.ToUpper(s)

	// save generated code to db
	result, err := codeCollection.InsertOne(context.TODO(), bson.M{
		"email": email,
		"code":  code,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	return code
}

func GetVerificationCode(code string) Verification {
	codeStruct := Verification{}

	// query
	query := bson.M{"code": code}
	if err2 := codeCollection.FindOne(context.TODO(), query).Decode(&codeStruct); err2 != nil {
		panic(err2)
	}
	jsonData, err := json.MarshalIndent(codeStruct, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	fmt.Println("Test123123")
	fmt.Println(codeStruct)
	return codeStruct
}
