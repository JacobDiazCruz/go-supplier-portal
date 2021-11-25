package person

import (
	"context"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/person/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var personCollection *mongo.Collection = database.OpenCollection(database.Client, "newsletter")

func SaveUser(person entity.Employee) string {
	// query
	result, err := personCollection.InsertOne(context.TODO(), person)
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	// test := map[string]string{
	// 	"id": oid.Hex(),
	// }
	// return get.GetUser(oid.Hex())
	return oid.Hex()
}
