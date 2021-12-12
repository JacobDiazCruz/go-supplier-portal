package contents

import (
	"context"
	"fmt"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/contents/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateService(content *entity.Content) string {
	// fmt.Println(content)
	// fmt.Println("here123123")
	// convert id string to mongo
	// objID, err := primitive.ObjectIDFromHex(contentId)
	// if err != nil {
	// 	panic(err)
	// }

	// query
	result, err := contentCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": content.ID},
		bson.M{
			"$set": bson.M{
				"title":          content.Title,
				"slug":           content.Slug,
				"tags":           content.Tags,
				"category":       content.Category,
				"status":         content.Status,
				"marketing_link": content.MarketingLink,
				"comments":       content.Comments,
				"vote_ids":       content.VoteIds,
				"total_votes":    content.TotalVotes,
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

func UpdateVoteIds(content *entity.ContentUpdates) string {
	// @todo: if vote_id exists already in the vote_ids, invalidate the vote

	// query
	result, err := contentCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": content.ID},
		bson.M{
			"$push": bson.M{
				"vote_ids": content.VoteId,
			},
			"$inc": bson.M{
				"total_votes": content.VoteAverage,
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
