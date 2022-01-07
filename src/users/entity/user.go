package users

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Email          string             `json:"email"`
	FirstName      string             `json:"firstname"`
	ThumbnailImage string             `json:"thumbnail_image"`
	OriginalImage  string             `json:"original_image"`
	LastName       string             `json:"lastname"`
	Password       string             `json:"password"`
	Role           string             `json:"role"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenIdentity struct {
	Token string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
