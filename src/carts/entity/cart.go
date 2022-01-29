package carts

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
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
	Products []ProductResponse  `json:"products" bson:"products"`
	UserId   primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	AuditLog AuditLog           `json:"audit_log"`
}

type AddToCart struct {
	ProductId string             `json:"product_id"`
	Quantity  int                `json:"quantity"`
	Variants  []VariantRequest   `json:"variants" bson:"variants"`
	UserId    primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	AuditLog  AuditLog           `json:"audit_log"`
}

type ProductResponse struct {
	ID                  primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name                string             `json:"name" validate:"required"`
	Quantity            float64            `json:"quantity"`
	Status              string             `json:"status" validate:"required"`
	Slug                string             `json:"slug" validate:"required"`
	Description         string             `json:"description" validate:"required"`
	Tags                []string           `json:"tags"`
	Category            string             `json:"category" validate:"required"`
	SalesInformation    SalesInformation   `json:"sales_information" bson:"sales_information" validate:"required"`
	Variants            []Variant          `json:"variants" bson:"variants"`
	TotalRatings        int                `json:"total_ratings" bson:"total_ratings" validate:"required"`
	ThumbnailDisplayUrl string             `json:"thumbnail_display_url" bson:"thumbnail_display_url"`
	Files               []File             `json:"files" bson:"files"`
	MarketingLink       string             `json:"marketing_link" bson:"marketing_link"`
	Reviews             []string           `json:"reviews"`
	AuditLog            AuditLog           `json:"audit_log" bson:"audit_log"`
}

type File struct {
	ThumbnailUrl string `json:"thumbnail_url" bson:"thumbnail_url"`
	OriginalUrl  string `json:"original_url" bson:"original_url"`
	FileType     string `json:"file_type" bson:"file_type"`
}

type SalesInformation struct {
	Price       float32 `json:"price" bson:"price" validate:"required"`
	SalePrice   float32 `json:"sale_price" bson:"sale_price"`
	Stock       float32 `json:"stock" bson:"stock" validate:"required"`
	Brand       string  `json:"brand" bson:"brand"`
	Sku         string  `json:"sku" bson:"sku" validate:"required"`
	MinQuantity float32 `json:"min_quantity" bson:"min_quantity" validate:"required"`
	MaxQuantity float32 `json:"max_quantity" bson:"max_quantity" validate:"required"`
}
type VariantRequest struct {
	VariantId       string `json:"variant_id" bson:"variant_id"`
	VariantOptionId string `json:"variant_option_id" bson:"variant_option_id"`
}

type Variant struct {
	ID     string      `json:"_id" bson:"_id"`
	Name   string      `json:"name" bson:"name"`
	Option interface{} `json:"option" bson:"option"`
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
