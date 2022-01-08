package carts

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	OrderStatus     OrderStatus        `json:"order_status"`
	DeliveryAddress string             `json:"delivery_address"`
	Cart            Cart               `json:"cart"`
	AuditLog        AuditLog           `json:"audit_log"`
}

type OrderStatus struct {
	Title string `json:"title"`
	Label string `json:"label"`
}

type PlaceOrder struct {
	CartId          string             `json:"cart_id"`
	OrderStatus     OrderStatus        `json:"order_status"`
	UserId          primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	DeliveryAddress string             `json:"delivery_address"`
	Note            string             `json:"note"`
	ShippingCourier string             `json:"shipping_courier"`
	ShippingAmount  int                `json:"shipping_amount"`
	AuditLog        AuditLog           `json:"audit_log"`
}

type SellerOrder struct {
	OrderId         primitive.ObjectID `json:"order_id"`
	Product         Product            `json:"product"`
	OrderStatus     OrderStatus        `json:"order_status"`
	SellerId        primitive.ObjectID `json:"order_id"`
	DeliveryAddress string             `json:"delivery_address"`
	Quantity        int                `json:"quantity"`
	AuditLog        AuditLog           `json:"audit_log"`
}

type Cart struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Products []Product          `json:"products"`
	UserId   primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	AuditLog AuditLog           `json:"audit_log"`
}

type Product struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name             string             `json:"name"`
	Status           string             `json:"status"`
	Slug             string             `json:"slug"`
	Quantity         float64            `json:"quantity"`
	Description      string             `json:"description"`
	Tags             []string           `json:"tags"`
	Category         string             `json:"category"`
	SalesInformation SalesInformation   `json:"sales_information"`
	Variation        []Variation        `json:"variation"`
	ThumbnailImage   string             `json:"thumbnail_image"`
	OriginalImage    string             `json:"original_image"`
	MarketingLink    string             `json:"marketing_link"`
	Reviews          []string           `json:"reviews"`
	AuditLog         AuditLog           `json:"audit_log"`
}

type Specification struct {
	Brand           string `json:"brand"`
	CountryOfOrigin string `json:"country_of_origin"`
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

type Auth struct {
	UserId primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
