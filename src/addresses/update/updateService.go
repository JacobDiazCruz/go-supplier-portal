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
	// id to mongoId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	// query
	result, err := addressCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objId},
		bson.M{
			"$push": bson.M{
				"addresses": bson.M{
					"profile_id":   address.ProfileId,
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
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	// return
	return id
}
