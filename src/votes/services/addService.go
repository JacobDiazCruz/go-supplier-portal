package votes

import (
	"context"

	contentService "gitlab.com/JacobDCruz/supplier-portal/src/contents/services"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/votes/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddService(vote entity.Vote, contentId string) string {
	// add other requests
	vote.ContentId = contentId

	// compute average vote
	totalAverage := vote.Creativity + vote.Graphics + vote.StoryTelling + vote.Impact
	vote.Average = totalAverage

	// query
	result, err := voteCollection.InsertOne(context.TODO(), bson.M{
		"content_id":   vote.ContentId,
		"creativity":   vote.Creativity,
		"graphics":     vote.Graphics,
		"storytelling": vote.StoryTelling,
		"impact":       vote.Impact,
		"average":      vote.Average,
		"audit_log":    vote.AuditLog,
	})
	if err != nil {
		panic(err)
	}
	oid := result.InsertedID.(primitive.ObjectID)

	// insert content_id to contents service
	contentService.UpdateService(oid.Hex(), contentId)

	// return
	return oid.Hex()
}
