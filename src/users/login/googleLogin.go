package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "gitlab.com/JacobDCruz/supplier-portal/src/auth"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/users/entity"
	get "gitlab.com/JacobDCruz/supplier-portal/src/users/get"
	signup "gitlab.com/JacobDCruz/supplier-portal/src/users/signup"
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

func GoogleLogin(ctx *gin.Context) {
	credentials := &auth.Credentials{}
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
		Email: credentials.Email,
	}

	// validate email if already exist
	emailRes, err := get.GetEmail(user.Email)
	if err != nil {
		fmt.Println("Err")
	}

	// Signup email if it doesnt exist yet
	if emailRes.Email == "" {
		res := signup.SignupService(user)
		fmt.Println(res)
		fmt.Println("here signup already")
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
		fmt.Println(tk)
		signToken := auth.SignToken(tk.Username)
		return signToken
	}
	return "Error"
}
