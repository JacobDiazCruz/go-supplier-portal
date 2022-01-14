package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type TokenIdentity struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	ThumbnailImage string `json:"thumbnail_image"`
	OriginalImage  string `json:"original_image"`
	Token          string `json:"token"`
}
type Claims struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	ThumbnailImage string `json:"thumbnail_image"`
	OriginalImage  string `json:"original_image"`
	jwt.RegisteredClaims
}

type Credentials struct {
	Email          string `json:"email" validate:"required"`
	Username       string `json:"username"`
	ThumbnailImage string `json:"thumbnail_image"`
	OriginalImage  string `json:"original_image"`
	Role           string `json:"role"`
	Token          string `json:"token"`
	Type           string `json:"type"`
	Scope          string `json:"scope"`
}
