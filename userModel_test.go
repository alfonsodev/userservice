package userservice

import (
	"code.google.com/p/google-api-go-client/plus/v1"
	//	"fmt"
	//	ara "github.com/diegogub/aranGO"
	"testing"
)

// func TestNew(t *testing.T) {
// 	//	var user UserModel
// 	_, err := ara.Connect("http://localhost:8529", "", "", false)
// 	if err != nil {
// 		t.Error("There is an error , can't connect")
// 	}
// 	//user.New(s)
// }

func TestSave(t *testing.T) {
	var person plus.Person
	user := NewUserModel()
	u2 := NewUserModel()
	person.Language = "en"
	u2.SetUserName("Ximo")
	user.SetUserName("Miguel")
	user.Doc.Person = person
	u2.Doc.Person = person
	user.save()
	u2.save()
}

func TestGetByGoogleId(t *testing.T) {
	user, err := GetByGoogleId("115601102326911748945")
	if err != nil {
		panic(err)
	}

	if user.Doc.Key == "" {
		t.Error("Cant get user by google id")
	}
}
