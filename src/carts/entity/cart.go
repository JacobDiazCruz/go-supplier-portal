package carts

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Products []string           `json:"products" bson:"products"`
	UserId   primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	AuditLog AuditLog           `json:"audit_log"`
}

type GetCart struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Products []bson.M           `json:"products"`
	UserId   primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	AuditLog AuditLog           `json:"audit_log"`
}
type ProductItems struct {
	ProductId bson.M  `json:"product_id"`
	Quantity  float32 `json:"quantity"`
}

type AddToCart struct {
	ProductId string             `json:"product_id"`
	Quantity  float32            `json:"quantity"`
	Variants  []VariantRequest   `json:"variants" bson:"variants"`
	UserId    primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	AuditLog  AuditLog           `json:"audit_log"`
}

type ProductResponse struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Quantity         float32            `json:"quantity"`
	Name             string             `json:"name"`
	Status           string             `json:"status"`
	ThumbnailImage   string             `json:"thumbnail_image"`
	OriginalImage    string             `json:"original_image"`
	Variants         []Variant          `json:"variants"`
	SalesInformation SalesInformation   `json:"sales_information"`
}

type SalesInformation struct {
	Price       float32 `json:"price"`
	Stock       float32 `json:"stock"`
	Brand       string  `json:"brand"`
	MinQuantity float32 `json:"min_quantity"`
	MaxQuantity float32 `json:"max_quantity"`
	UnitPrice   float32 `json:"unit_price"`
}

type VariantRequest struct {
	VariantId       string `json:"variant_id" bson:"variant_id"`
	VariantOptionId string `json:"variant_option_id" bson:"variant_option_id"`
}

type Variant struct {
	ID     string      `json:"_id" bson:"_id"`
	Name   string      `json:"name" bson:"name"`
	Option interface{} `json:"options" bson:"options"`
}

type VariantOptions struct {
	Name           string  `json:"name" bson:"name"`
	Price          float32 `json:"price" bson:"price"`
	Stock          float32 `json:"stock" bson:"stock"`
	Sku            string  `json:"sku" bson:"sku"`
	SalePrice      float32 `json:"sale_price" bson:"sale_price"`
	MinQuantity    float32 `json:"min_quantity" bson:"min_quantity"`
	MaxQuantity    float32 `json:"max_quantity" bson:"max_quantity"`
	ThumbnailImage string  `json:"thumbnail_image" bson:"thumbnail_image"`
	OriginalImage  string  `json:"original_image" bson:"original_image"`
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
