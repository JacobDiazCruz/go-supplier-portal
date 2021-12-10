package votes

import (
	"context"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/votes/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddService(vote entity.Vote, nomineeId string) string {
	// add other requests
	vote.NomineeId = nomineeId

	// compute average vote
	totalAverage := vote.Creativity + vote.Graphics + vote.StoryTelling + vote.Impact
	vote.Average = totalAverage

	// query
	result, err := voteCollection.InsertOne(context.TODO(), vote)
	if err != nil {
		panic(err)
	}

	// return
	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
