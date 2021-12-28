package carts

import (
	"context"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "carts")

// 1. create a cart per user
// 2. push item and its quantity per user's cart
func AddService(cart entity.Cart) string {
	// query
	result, err := cartCollection.InsertOne(context.TODO(), bson.M{
		"products":  cart.Products,
		"user_id":   cart.UserId,
		"audit_log": cart.AuditLog,
	})
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
