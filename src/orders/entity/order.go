package carts

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type PlaceOrder struct {
	CartId          string   `json:"cart_id"`
	DeliveryAddress string   `json:"delivery_address"`
	AuditLog        AuditLog `json:"audit_log"`
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
