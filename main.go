package userservice

import (
	"fmt"
	"net/http"

	googleAuth "github.com/alfonsodev/googleauth"
	"github.com/flosch/pongo2"
	//  "github.com/gorilla/sessions"
)

func main() {
	fmt.Printf("\nHello user service...\n>>>>>\n")
}

// var store = sessions.NewCookieStore([]byte("1--!!@323kjlkb1#@$V3k1jb31S}{23jcl2"))

func AuthHandler(rw http.ResponseWriter, r *http.Request) {
	//  session, _ := store.Get(r, "session-name")
	//  fmt.Printf("session: %v", session.Values["foo"])
	//  session.Values[42] = 43
	// Save it.
	//  session.Save(r, rw)

	//  fmt.Printf("session: %+v", session.Values)
	url := googleAuth.GetGoogleAuthUrl()
	tmpl := pongo2.Must(pongo2.FromFile("templates/dashboard.html"))
	err := tmpl.ExecuteWriter(pongo2.Context{"GoogleAuthUrl": url}, rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
