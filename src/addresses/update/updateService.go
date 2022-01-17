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
				"user_id":  address.UserId,
				"fullname": address.FullName,
				"region": bson.M{
					"region_name": address.Region.RegionName,
					"region_code": address.Region.RegionCode,
					"psgc_code":   address.Region.PsgcCode,
				},
				"province": bson.M{
					"province_name": address.Province.ProvinceName,
					"province_code": address.Province.ProvinceCode,
					"psgc_code":     address.Province.PsgcCode,
				},
				"city": bson.M{
					"city_name":     address.City.CityName,
					"city_code":     address.City.CityCode,
					"region_code":   address.City.RegionCode,
					"province_code": address.Province.ProvinceCode,
				},
				"barangay": bson.M{
					"brgy_name":     address.Barangay.BrgyName,
					"brgy_code":     address.Barangay.BrgyCode,
					"region_code":   address.Barangay.RegionCode,
					"province_code": address.City.ProvinceCode,
				},
				"street":       address.Street,
				"label":        address.Label,
				"phone_number": address.PhoneNumber,
				"postal_code":  address.PostalCode,
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
