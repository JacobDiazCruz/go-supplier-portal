package reviews

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ProductId primitive.ObjectID `json:"product_id" bson:"product_id"`
	Rating    int                `json:"rating" bson:"rating"`
	Comment   string             `json:"comment" bson:"comment"`
	AuditLog  AuditLog           `json:"audit_log" bson:"audit_log"`
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

type Auth struct {
	UserId primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
