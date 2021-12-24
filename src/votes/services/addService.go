package votes

import (
	"context"

	contentEntity "gitlab.com/JacobDCruz/supplier-portal/src/contents/entity"
	contentService "gitlab.com/JacobDCruz/supplier-portal/src/contents/services"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/votes/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddService(vote *entity.Vote, contentId string, audit *entity.AuditLog) string {
	// add contentId
	objContentId, err := primitive.ObjectIDFromHex(contentId)
	if err != nil {
		panic(err)
	}
	vote.ContentId = objContentId

	// compute average vote
	totalAverage := vote.Creativity + vote.Graphics + vote.StoryTelling + vote.Impact
	vote.Average = totalAverage

	// @TODO: add audit log
	// audit := &entity.AuditLog{
	// 	Name:      ct.Username,
	// 	CreatedAt: time.Now(),
	// 	CreatedBy: ct.Username,
	// 	UpdatedAt: time.Now(),
	// 	UpdatedBy: "",
	// }
	// fmt.Println(audit)
	// fmt.Println("trashtasrqwrqwr")

	// query
	result, err := voteCollection.InsertOne(context.TODO(), bson.M{
		"content_id":   objContentId,
		"creativity":   vote.Creativity,
		"graphics":     vote.Graphics,
		"storytelling": vote.StoryTelling,
		"impact":       vote.Impact,
		"average":      vote.Average,
		"audit_log":    audit,
	})
	if err != nil {
		panic(err)
	}
	oid := result.InsertedID.(primitive.ObjectID)

	// insert content_id to contents service
	params := &contentEntity.ContentUpdates{}
	objID, err := primitive.ObjectIDFromHex(contentId)
	if err != nil {
		panic(err)
	}
	params.ID = objID
	params.VoteId = oid.Hex()
	params.VoteAverage = vote.Average
	contentService.UpdateVoteIds(params)

	// return string _id
	return oid.Hex()
}
