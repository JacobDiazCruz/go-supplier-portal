package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	createCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/create"
	cartEntity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	get "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
	signup "gitlab.com/JacobDCruz/supplier-portal/src/users/signup"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/oauth2/v2"
)

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
// var googleOauthConfig = &oauth2.Config{
// 	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
// 	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
// 	RedirectURL:  "http://localhost:300/googlelogin",
// 	Scopes: []string{
// 		"https://www.googleapis.com/auth/userinfo.profile",
// 		"https://www.googleapis.com/auth/userinfo.email",
// 	},
// 	Endpoint: google.Endpoint,
// }

// const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

var httpClient = &http.Client{}

var credentials = &auth.Credentials{}

func GoogleLogin(ctx *gin.Context) {
	err2 := ctx.BindJSON(&credentials)
	fmt.Println("im here 0")
	if err2 != nil {
		fmt.Println("im here 1")
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}
	token := verifyIdToken(credentials.Token)
	// @TODO: Signup google email if it does not exist in db
	// bind requestData
	user := entity.User{
		Email:          credentials.Email,
		Username:       credentials.Username,
		ThumbnailImage: credentials.ThumbnailImage,
		OriginalImage:  credentials.OriginalImage,
		Role:           credentials.Role,
	}

	// validate email if already exist
	emailRes := get.GetEmail(user.Email)

	// Signup email if it doesnt exist yet
	if emailRes.Email == "" {
		res := signup.SignupService(user)
		objID, err := primitive.ObjectIDFromHex(res)
		if err != nil {
			panic(err)
		}
		// Create a cart for the new signed up user
		cEntity := &cartEntity.Cart{}
		cEntity.UserId = objID
		cEntity.Products = []string{}
		createCart.AddService(*cEntity)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": token})
}

// Verify Id Token func
func verifyIdToken(idToken string) string {
	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return "Error"
	}

	// if token is valid, return token
	if tokenInfo.VerifiedEmail == true {
		tk := &auth.TokenIdentity{}
		tk.Email = credentials.Email
		signToken := auth.SignToken(tk.Email)
		return signToken
	}
	return "Error"
}
