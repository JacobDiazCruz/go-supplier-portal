package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func verifyIdToken(idToken string) string {
	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return "Error"
	}
	fmt.Println(tokenInfo)
	fmt.Println("testest")
	if tokenInfo.VerifiedEmail == true {
		tk := &TokenIdentity{}
		fmt.Println(tk)
		signToken := SignToken(tk.Username)
		return signToken
	}
	return "Here"
}

func GoogleLogin(ctx *gin.Context) {
	credentials := &Credentials{}
	err2 := ctx.BindJSON(&credentials)
	fmt.Println("im here 0")
	if err2 != nil {
		fmt.Println("im here 1")
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Error encountered"})
	}
	token := verifyIdToken(credentials.Token)

	// @TODO: Signup google email if it does not exist in db

	ctx.JSON(http.StatusOK, gin.H{"data": token})
	// client, err := google.DefaultClient(context.Background(), credentials.Scope)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client.Get("")
	// fmt.Println(client)
	// fmt.Println("24242424242")

	// var token string          // this comes from your web or mobile app maybe
	// googleClientId := os.Getenv("GOOGLE_OAUTH_CLIENT_ID") // from credentials in the Google dev console

	// tokenValidator, err := idtoken.NewValidator(context.Background())
	// if err != nil {
	// 	fmt.Println("im here 2")
	// 	// handle error, stop execution
	// }

	// fmt.Println(tokenValidator)
	// fmt.Println("here payload1231231232")
	// payload, err := tokenValidator.Validate(context.Background(), credentials.Token, googleClientId)
	// if err != nil {
	// 	// fmt.Println(payload)
	// 	fmt.Println("im here 3")
	// }
	// // // handle error, stop execution
	// email := payload.Claims["email"]
	// name := payload.Claims["name"]
	// fmt.Println(email)
	// fmt.Println(name)
}
