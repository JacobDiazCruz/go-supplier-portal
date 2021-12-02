package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtRefreshKey = []byte("secret_key")

func RefreshController(ctx *gin.Context) {
	cookie, err := ctx.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Bad Request"})
		return
	}

	tokenStr := cookie

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtRefreshKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Bad Request"})
		return
	}
	if !tkn.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Bad Request"})
	}

	expirationTime := time.Now().Add(time.Minute * 15)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtRefreshKey)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Error"})
		return
	}

	// http.SetCookie(w,
	// 	&http.Cookie{
	// 		Name:    "refresh_token",
	// 		Value:   tokenString,
	// 		Expires: expirationTime,
	// 	})

	ctx.SetCookie("refresh_token", tokenString, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Refresh Token Successfully", "data": tokenString})
}
