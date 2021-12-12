package contents

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Content struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title          string             `json:"title"`
	Slug           string             `json:"slug"`
	Body           string             `json:"body"`
	Tags           []string           `json:"tags"`
	Category       string             `json:"category"`
	ThumbnailImage string             `json:"thumbnail_image"`
	OriginalImage  string             `json:"original_image"`
	Status         string             `json:"status"`
	MarketingLink  string             `json:"marketing_link"`
	Comments       []string           `json:"comments"`
	VoteIds        []string           `json:"vote_ids"`
	TotalVotes     float64            `json:"total_votes"`
	AuditLog       AuditLog           `json:"audit_log"`
}

type ContentUpdates struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CommentId   string             `json:"comment_id"`
	VoteId      string             `json:"vote_id"`
	VoteAverage float64            `json:"vote_average"`
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
