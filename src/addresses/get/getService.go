package addresses

import (
	"context"
	"encoding/json"
	"fmt"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/addresses/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var addressCollection *mongo.Collection = database.OpenCollection(database.Client, "addresses")

type getService interface {
	GetService() entity.Address
}

type Param struct {
	UserId primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
}

func GetService(id string) entity.Address {
	// set initial values
	address := entity.Address{}
	var result bson.M

	var query = bson.M{"_id": ""}
	objID, err := primitive.ObjectIDFromHex(id)

	// profile_id query params
	query = bson.M{"_id": objID}
	fmt.Println(query)

	// query to db
	err2 := addressCollection.FindOne(context.TODO(), query).Decode(&result)
	if err2 != nil {
		panic(err2)
	}

	// unmarshal result to products struct
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &address)

	return address
}
