package addresses

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Region      Region             `json:"region"`
	Province    Province           `json:"province"`
	City        City               `json:"city"`
	Barangay    Barangay           `json:"barangay"`
	Street      string             `json:"street"`
	PostalCode  string             `json:"postal_code"`
	PhoneNumber string             `json:"phone_number"`
	FullName    string             `json:"fullname"`
	Label       string             `json:"label"`
	Default     bool               `json:"default"`
	AuditLog    AuditLog           `json:"audit_log"`
}

type Region struct {
	RegionName string `json:"region_name"`
	RegionCode string `json:"region_code"`
	PsgcCode   string `json:"psgc_code"`
}
type Province struct {
	ProvinceName string `json:"province_name"`
	ProvinceCode string `json:"province_code"`
	RegionCode   string `json:"region_code"`
	PsgcCode     string `json:"psgc_code"`
}
type City struct {
	CityName     string `json:"city_name"`
	CityCode     string `json:"city_code"`
	ProvinceCode string `json:"province_code"`
	RegionCode   string `json:"region_code"`
}
type Barangay struct {
	BrgyName     string `json:"brgy_name"`
	BrgyCode     string `json:"brgy_code"`
	ProvinceCode string `json:"province_code"`
	RegionCode   string `json:"region_code"`
}

type AuditLog struct {
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	ThumbnailImage string    `json:"thumbnail_image"`
	OriginalImage  string    `json:"original_image"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedBy      string    `json:"created_by"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedBy      string    `json:"updated_by"`
}

type TokenIdentity struct {
	Token string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
