package person

import (
	"context"
	"encoding/json"
	"fmt"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/person/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var personCollection *mongo.Collection = database.OpenCollection(database.Client, "newsletter")

type getService interface {
	GetDataById() entity.Employee
}

type Param struct {
	id string
}

// func GetService(g getService) entity.Employee {
// 	return g.GetDataById()
// }

func GetDataById(id string) entity.Employee {
	result := entity.Employee{}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	// query
	query := bson.M{"_id": objID}
	if err2 := personCollection.FindOne(context.TODO(), query).Decode(&result); err != nil {
		panic(err2)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	fmt.Println("Test123123")
	fmt.Println(result)
	return result
}
