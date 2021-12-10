package votes

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vote struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NomineeId    string             `json:"nominee_id"`
	Creativity   float64            `json:"creativity"`
	Graphics     float64            `json:"graphics"`
	StoryTelling float64            `json:"storytelling"`
	Impact       float64            `json:"impact"`
	Average      float64            `json:"average"`
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
