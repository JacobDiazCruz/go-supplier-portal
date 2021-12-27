package products

import (
	"context"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/products/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

func UpdateService(product entity.Product) string {
	// query
	result, err := productCollection.InsertOne(context.TODO(), bson.M{
		"title":           product.Name,
		"slug":            product.Slug,
		"description":     product.Description,
		"tags":            product.Tags,
		"category":        product.Category,
		"thumbnail_image": product.ThumbnailImage,
		"original_image":  product.OriginalImage,
		"status":          product.Status,
		"marketing_link":  product.MarketingLink,
		"comments":        product.Comments,
		"audit_log":       product.AuditLog,
	})
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
