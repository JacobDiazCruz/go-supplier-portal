package users

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Email          string             `json:"email" validate:"required"`
	Username       string             `json:"username" validate:"required"`
	ThumbnailImage string             `json:"thumbnail_image"`
	OriginalImage  string             `json:"original_image"`
	Password       string             `json:"password" validate:"required"`
	ContactNumber  int                `json:"contact_number" validate:"required"`
	Role           string             `json:"role"`
}

type Credentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type GoogleLoginRequest struct {
	Email          string `json:"email" validate:"required"`
	Username       string `json:"username" validate:"required"`
	ThumbnailImage string `json:"thumbnail_image"`
	OriginalImage  string `json:"original_image"`
	Role           string `json:"role"`
	Token          string `json:"token"`
	Type           string `json:"type"`
	Scope          string `json:"scope"`
}

type TokenIdentity struct {
	Token string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
