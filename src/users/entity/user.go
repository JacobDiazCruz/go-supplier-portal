package users

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Email          string             `json:"email" validate:"required"`
	FirstName      string             `json:"firstname" validate:"required"`
	ThumbnailImage string             `json:"thumbnail_image"`
	OriginalImage  string             `json:"original_image"`
	LastName       string             `json:"lastname" validate:"required"`
	Password       string             `json:"password" validate:"required"`
	Role           string             `json:"role"`
}

type Credentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenIdentity struct {
	Token string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
