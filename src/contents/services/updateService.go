package contents

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateService(voteId string, contentId string) string {
	fmt.Println(voteId)
	fmt.Println("232323232323232")

	// convert id string to mongo
	objID, err := primitive.ObjectIDFromHex(contentId)
	if err != nil {
		panic(err)
	}

	// query
	result, err := contentCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{
			"$push": bson.M{
				"vote_ids": voteId,
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("update service here id ^")
	return "Success"
}
