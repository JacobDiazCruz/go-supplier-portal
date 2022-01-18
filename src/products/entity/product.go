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
	SalesInformation SalesInformation   `json:"sales_information" bson:"sales_information" validate:"required"`
	Variation        []Variation        `json:"variation" validate:"required"`
	TotalRatings     int                `json:"total_ratings" bson:"total_ratings" validate:"required"`
	ThumbnailImage   string             `json:"thumbnail_image" bson:"thumbnail_image" validate:"required"`
	OriginalImage    string             `json:"original_image" bson:"original_image" validate:"required"`
	MarketingLink    string             `json:"marketing_link"`
	Reviews          []string           `json:"reviews"`
	AuditLog         AuditLog           `json:"audit_log"`
}

type Specification struct {
	Brand           string `json:"brand"`
	CountryOfOrigin string `json:"country_of_origin"`
}

type List struct {
	Order   string `json:"order"`
	OrderBy string `json:"order_by" bson:"order_by"`
	PageNum int    `json:"page_num" bson:"order_by"`
	Limit   int64  `json:"limit"`
	Sort    string `json:"sort"`
}

type SalesInformation struct {
	Price       float32 `json:"price" validate:"required"`
	Stock       float32 `json:"stock" validate:"required"`
	Brand       string  `json:"brand"`
	MinQuantity float32 `json:"min_quantity" bson:"min_quantity" validate:"required"`
	MaxQuantity float32 `json:"max_quantity" bson:"max_quantity" validate:"required"`
	UnitPrice   float32 `json:"unit_price" bson:"unit_price"`
}

type Variation struct {
	Name           string  `json:"name"`
	Price          float32 `json:"price"`
	Stock          float32 `json:"stock"`
	MinQuantity    float32 `json:"min_quantity" bson:"min_quantity"`
	MaxQuantity    float32 `json:"max_quantity" bson:"max_quantity"`
	DiscountPrice  float32 `json:"discount_price" bson:"discount_price"`
	ThumbnailImage string  `json:"thumbnail_image" bson:"thumbnail_image"`
	OriginalImage  string  `json:"original_image" bson:"original_image"`
}

type Others struct {
	PreOrder  bool   `json:"pre_order"`
	Condition bool   `json:"condition"`
	ParentSKU string `json:"parent_sku"`
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
