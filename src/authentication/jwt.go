package authentication

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(ctx *gin.Context) {
	fmt.Println("Im here")
	var credentials Credentials
	err := ctx.BindJSON(&credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Bad Req"})
		return
	}

	expectedPassword, ok := users[credentials.Username]

	if !ok || expectedPassword != credentials.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauth"})
		return
	}
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	fmt.Println(tokenString)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Server Error"})
		return
	}

	ctx.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"msg": "Sucess Login", "data": tokenString})
}

// func Home(w http.ResponseWriter, r *http.Request) {
// 	cookie, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	tokenStr := cookie.Value

// 	claims := &Claims{}

// 	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
// 		func(t *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil
// 		})

// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))

// }

func Refresh(ctx *gin.Context) {
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
			return jwtKey, nil
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

	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(time.Minute * 15)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

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
