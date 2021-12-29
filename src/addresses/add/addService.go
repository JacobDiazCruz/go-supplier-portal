package addresses

import (
	"context"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/addresses/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var addressesCollection *mongo.Collection = database.OpenCollection(database.Client, "addresses")

func AddService(address entity.Address) string {
	// query
	result, err := addressesCollection.InsertOne(context.TODO(), bson.M{
		"user_id":      address.UserId,
		"region":       address.Region,
		"province":     address.Province,
		"city":         address.City,
		"barangay":     address.Barangay,
		"label":        address.Label,
		"phone_number": address.PhoneNumber,
		"default":      address.Default,
		"audit_log":    address.AuditLog,
	})
	if err != nil {
		panic(err)
	}

	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
