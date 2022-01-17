package orders

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	OrderStatus     OrderStatus        `json:"order_status" bson:"order_status"`
	DeliveryAddress Address            `json:"delivery_address" bson:"delivery_address"`
	Cart            Cart               `json:"cart"`
	Note            string             `json:"note"`
	PaymentMethod   string             `json:"payment_method" bson:"payment_method"`
	ShippingCourier string             `json:"shipping_courier" bson:"shipping_courier"`
	ShippingAmount  int                `json:"shipping_amount" bson:"shipping_amount"`
	SubtotalAmount  int                `json:"subtotal_amount" bson:"subtotal_amount"`
	TotalAmount     int                `json:"total_amount" bson:"total_amount"`
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
	DeliveryAddress Address            `json:"delivery_address" bson:"delivery_address"`
	Note            string             `json:"note" bson:"note"`
	PaymentMethod   string             `json:"payment_method" bson:"payment_method"`
	ShippingCourier string             `json:"shipping_courier" bson:"shipping_courier"`
	ShippingAmount  int                `json:"shipping_amount" bson:"shipping_amount"`
	AuditLog        AuditLog           `json:"audit_log" bson:"audit_log"`
}

type SellerOrder struct {
	OrderId         primitive.ObjectID `json:"order_id"`
	Product         Product            `json:"product"`
	OrderStatus     OrderStatus        `json:"order_status"`
	SellerId        primitive.ObjectID `json:"order_id"`
	DeliveryAddress Address            `json:"delivery_address"`
	Quantity        int                `json:"quantity"`
	AuditLog        AuditLog           `json:"audit_log"`
}

type Cart struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Products []Product          `json:"products"`
	UserId   primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	AuditLog AuditLog           `json:"audit_log"`
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
