package products

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	ProductId primitive.ObjectID `json:"product_id" bson:"product_id,omitempty"`
	Quantity  float32            `json:"quantity"`
	AuditLog  AuditLog           `json:"audit_log"`
}

type Product struct {
	Name             string           `json:"name"`
	Status           string           `json:"status"`
	ThumbnailImage   string           `json:"thumbnail_image"`
	OriginalImage    string           `json:"original_image"`
	Variation        []Variation      `json:"variation"`
	SalesInformation SalesInformation `json:"sales_information"`
}

type SalesInformation struct {
	Price       float32 `json:"price"`
	Stock       float32 `json:"stock"`
	Brand       string  `json:"brand"`
	MinQuantity float32 `json:"min_quantity"`
	MaxQuantity float32 `json:"max_quantity"`
	UnitPrice   float32 `json:"unit_price"`
}

type Variation struct {
	Name          string  `json:"name"`
	Price         float32 `json:"price"`
	Stock         float32 `json:"stock"`
	MinQuantity   float32 `json:"min_quantity"`
	MaxQuantity   float32 `json:"max_quantity"`
	DiscountPrice float32 `json:"discount_price"`
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
