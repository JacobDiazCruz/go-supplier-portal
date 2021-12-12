package votes

import (
	"context"
	"encoding/json"
	"fmt"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/votes/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetService(id string) entity.Vote {
	result := entity.Vote{}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	// query
	query := bson.M{"_id": objID}
	if err2 := voteCollection.FindOne(context.TODO(), query).Decode(&result); err != nil {
		panic(err2)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	fmt.Println("Test123123")
	fmt.Println(result)
	// return
	return result
}
