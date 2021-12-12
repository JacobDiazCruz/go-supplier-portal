package users

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Password  string             `json:"password"`
	Role      string             `json:"role"`
	ProfileId string             `json:"profile_id"`
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
