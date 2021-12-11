package contents

import (
	"context"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/contents/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddService(content entity.Content) string {
	// query
	result, err := contentCollection.InsertOne(context.TODO(), bson.M{
		"title":           content.Title,
		"slug":            content.Slug,
		"body":            content.Body,
		"tags":            content.Tags,
		"category":        content.Category,
		"thumbnail_image": content.ThumbnailImage,
		"original_image":  content.OriginalImage,
		"status":          content.Status,
		"marketing_link":  content.MarketingLink,
		"comments":        content.Comments,
		"vote_ids":        content.VoteIds,
		"total_votes":     content.TotalVotes,
		"audit_log":       content.AuditLog,
	})
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
