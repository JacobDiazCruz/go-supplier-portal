package profiles

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Role      string             `json:"role"`
	UserId    primitive.ObjectID `json:"id" bson:"user_id,omitempty"`
}

type TokenIdentity struct {
	Token string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
