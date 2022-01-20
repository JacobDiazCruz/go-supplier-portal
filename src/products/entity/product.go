package products

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name             string             `json:"name" validate:"required"`
	Status           string             `json:"status" validate:"required"`
	Slug             string             `json:"slug" validate:"required"`
	Description      string             `json:"description" validate:"required"`
	Tags             []string           `json:"tags"`
	Category         string             `json:"category" validate:"required"`
	Wholesale        []Wholesale        `json:"wholesale" bson:"wholesale"`
	SalesInformation SalesInformation   `json:"sales_information" bson:"sales_information" validate:"required"`
	Variants         []Variant          `json:"variants" validate:"required"`
	TotalRatings     int                `json:"total_ratings" bson:"total_ratings" validate:"required"`
	ThumbnailImage   string             `json:"thumbnail_image" bson:"thumbnail_image" validate:"required"`
	OriginalImage    string             `json:"original_image" bson:"original_image" validate:"required"`
	MarketingLink    string             `json:"marketing_link" bson:"marketing_link"`
	Reviews          []string           `json:"reviews" bson:"reviews"`
	AuditLog         AuditLog           `json:"audit_log" bson:"audit_log"`
}

type Wholesale struct {
	MinPrice  string `json:"min_price" bson:"min_price"`
	MaxPrice  string `json:"max_price" bson:"max_price"`
	UnitPrice string `json:"unit_price" bson:"unit_price"`
}

type List struct {
	Order   string `json:"order"`
	OrderBy string `json:"order_by" bson:"order_by"`
	PageNum int    `json:"page_num" bson:"order_by"`
	Limit   int64  `json:"limit"`
	Sort    string `json:"sort"`
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

type Variant struct {
	Name    string           `json:"name" bson:"name"`
	Options []VariantOptions `json:"options" bson:"options"`
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

type ProductUpdates struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CommentId   string             `json:"comment_id"`
	VoteId      string             `json:"vote_id"`
	VoteAverage float32            `json:"vote_average"`
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
