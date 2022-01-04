package reviews

import (
	"context"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/reviews/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var reviewCollection *mongo.Collection = database.OpenCollection(database.Client, "reviews")

func AddService(review entity.Review) string {
	// query
	result, err := reviewCollection.InsertOne(context.TODO(), bson.M{
		"product_id": review.ProductId,
		"rating":     review.Rating,
		"comment":    review.Comment,
		"audit_log":  review.AuditLog,
	})
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
