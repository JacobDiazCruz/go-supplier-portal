package addresses

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Region      Region             `json:"region" bson:"region" validate:"required"`
	Province    Province           `json:"province" bson:"province" validate:"required"`
	City        City               `json:"city" bson:"city" validate:"required"`
	Barangay    Barangay           `json:"barangay" bson:"barangay" validate:"required"`
	Street      string             `json:"street" bson:"street" validate:"required"`
	PostalCode  string             `json:"postal_code" bson:"postal_code" validate:"required"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number" validate:"required"`
	FullName    string             `json:"fullname" bson:"fullname" validate:"required"`
	Label       string             `json:"label" bson:"label"`
	Default     bool               `json:"default" bson:"default" validate:"required"`
	AuditLog    AuditLog           `json:"audit_log" bson:"audit_log" validate:"required"`
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

type AuditLog struct {
	Name           string    `json:"name" bson:"name"`
	Email          string    `json:"email" bson:"email"`
	ThumbnailImage string    `json:"thumbnail_image" bson:"thumbnail_image"`
	OriginalImage  string    `json:"original_image" bson:"original_image"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	CreatedBy      string    `json:"created_by" bson:"created_by"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy      string    `json:"updated_by" bson:"updated_by"`
}

type TokenIdentity struct {
	Token string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
