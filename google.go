package userservice

import (
	"code.google.com/p/goauth2/oauth"
	"code.google.com/p/google-api-go-client/plus/v1"
	"fmt"
	googleAuth "github.com/alfonsodev/googleauth"
)

var config = &oauth.Config{
	ClientId:     clientID,
	ClientSecret: clientSecret,
	// Scope determines which API calls you are authorized to make
	Scope:    "https://www.googleapis.com/auth/plus.login",
	AuthURL:  "https://accounts.google.com/o/oauth2/auth",
	TokenURL: "https://accounts.google.com/o/oauth2/token",
	//Use "postmessage" for the code-flow for server side apps
	RedirectURL: "postmessage",
}

// Get the token
// Save in the database
// Start a user session
// Redirect to user home (dashboard)
func GoogleAuthLogic(code string) bool {
	accessToken, idToken, err := googleAuth.Exchange(code)
	gplusID, err := googleAuth.DecodeIdToken(idToken)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n Gplus:%+v\n", gplusID)
	if gplusID != "115601102326911748945" {
		return false
	} else {
		user, _ := GetByGoogleId(gplusID)
		if user.Doc.Key == "" {
			person := getPersonFromToken(accessToken)
			user := NewUserModel()
			user.Doc.Person = *person
			user.save()
		}
		return true
	}
}

func getPersonFromToken(token string) (person *plus.Person) {
	// Create a new authorized API client
	t := &oauth.Transport{Config: config}
	tok := new(oauth.Token)
	tok.AccessToken = token
	t.Token = tok
	service, err := plus.New(t.Client())
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}
	// Get a list of people that this user has shared with this app
	people := service.People.Get("me")
	person, err = people.Do()
	//TODO:Handle all this posible errors
	if err != nil {
		if err.Error() == "AccessTokenRefreshError" {
			fmt.Printf("\n err: %s", err)
			return // &appError{errors.New(m), m, 500}
		}
		fmt.Printf("\n err: %s", err)
		return // &appError{err, m, 500}
	}

	return person
}
