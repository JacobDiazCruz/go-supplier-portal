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
		"title":                 product.Name,
		"slug":                  product.Slug,
		"description":           product.Description,
		"tags":                  product.Tags,
		"category":              product.Category,
		"files":                 product.Files,
		"status":                product.Status,
		"thumbnail_display_url": product.ThumbnailDisplayUrl,
		"marketing_link":        product.MarketingLink,
		"reviews":               product.Reviews,
		"audit_log":             product.AuditLog,
	})
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}

func UpdateStock(productId primitive.ObjectID, quantity float64) string {
	// query filters
	filter := bson.M{"_id": productId}
	update := bson.M{"sales_information.stock": -quantity}

	// query db
	res := productCollection.FindOneAndUpdate(context.Background(),
		filter,
		bson.M{"$inc": update},
	)

	// check error
	if res.Err() != nil {
		panic(res.Err())
	}

	return "Sucess"
}

func UpdateTotalRatings(productId primitive.ObjectID, count int) string {
	// query filters
	filter := bson.M{"_id": productId}
	update := bson.M{"total_ratings": +count}

	// query db
	res := productCollection.FindOneAndUpdate(context.Background(),
		filter,
		bson.M{"$inc": update},
	)

	// check error
	if res.Err() != nil {
		panic(res.Err())
	}

	return "Sucess"
}
