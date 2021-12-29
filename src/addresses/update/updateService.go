package addresses

import (
	"context"
	"fmt"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/addresses/entity"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var addressCollection *mongo.Collection = database.OpenCollection(database.Client, "addresses")

func UpdateService(address entity.Address, id string) string {
	// convert id string to mongo
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	// query
	result, err := addressCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{
			"$set": bson.M{
				"user_id":      address.UserId,
				"region":       address.Region,
				"province":     address.Province,
				"city":         address.City,
				"barangay":     address.Barangay,
				"label":        address.Label,
				"phone_number": address.PhoneNumber,
				"default":      address.Default,
				"audit_log":    address.AuditLog,
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	return "Success"
}
