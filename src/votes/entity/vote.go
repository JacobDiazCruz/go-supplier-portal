package votes

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vote struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ContentId    primitive.ObjectID `json:"content_id" bson:"content_id,omitempty"`
	Creativity   float32            `json:"creativity"`
	Graphics     float32            `json:"graphics"`
	StoryTelling float32            `json:"storytelling"`
	Impact       float32            `json:"impact"`
	Average      float32            `json:"average"`
	AuditLog     AuditLog           `json:"audit_log"`
}

type AuditLog struct {
	Name           string    `json:"name"`
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
