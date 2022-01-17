package orders

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Region      Region             `json:"region" bson:"region"`
	Province    Province           `json:"province" bson:"province"`
	City        City               `json:"city" bson:"city"`
	Barangay    Barangay           `json:"barangay" bson:"barangay"`
	Street      string             `json:"street" bson:"street"`
	PostalCode  string             `json:"postal_code" bson:"postal_code"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	FullName    string             `json:"fullname" bson:"fullname"`
	Label       string             `json:"label" bson:"label"`
	Default     bool               `json:"default" bson:"default"`
	AuditLog    AuditLog           `json:"audit_log" bson:"audit_log"`
}

type Region struct {
	RegionName string `json:"region_name" bson:"region_name"`
	RegionCode string `json:"region_code" bson:"region_code"`
	PsgcCode   string `json:"psgc_code" bson:"psgc_code"`
}
type Province struct {
	ProvinceName string `json:"province_name" bson:"province_name"`
	ProvinceCode string `json:"province_code" bson:"province_code"`
	RegionCode   string `json:"region_code" bson:"region_code"`
	PsgcCode     string `json:"psgc_code" bson:"psgc_code"`
}
type City struct {
	CityName     string `json:"city_name" bson:"city_name"`
	CityCode     string `json:"city_code" bson:"city_code"`
	ProvinceCode string `json:"province_code" bson:"province_code"`
	RegionCode   string `json:"region_code" bson:"region_code"`
}
type Barangay struct {
	BrgyName     string `json:"brgy_name" bson:"brgy_name"`
	BrgyCode     string `json:"brgy_code" bson:"brgy_code"`
	ProvinceCode string `json:"province_code" bson:"province_code"`
	RegionCode   string `json:"region_code" bson:"region_code"`
}
