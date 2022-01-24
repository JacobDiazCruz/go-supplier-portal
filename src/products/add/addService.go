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
var variantCollection *mongo.Collection = database.OpenCollection(database.Client, "variant_options")

func AddProductService(product entity.Product) string {
	// query
	result, err := productCollection.InsertOne(context.TODO(), bson.M{
		"name":                  product.Name,
		"slug":                  product.Slug,
		"body":                  product.Description,
		"tags":                  product.Tags,
		"category":              product.Category,
		"total_ratings":         product.TotalRatings,
		"wholesale":             product.Wholesale,
		"sales_information":     product.SalesInformation,
		"files":                 product.Files,
		"variants":              product.Variants,
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
