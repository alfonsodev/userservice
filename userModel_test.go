package userservice

import (
	"code.google.com/p/google-api-go-client/plus/v1"
	"fmt"
	ara "github.com/diegogub/aranGO"
	"testing"
)

func TestNew(t *testing.T) {
	//	var user UserModel
	_, err := ara.Connect("http://localhost:8529", "", "", false)
	if err != nil {
		t.Error("There is an error , can't connect")
	}
	//user.New(s)
}

func TestSave(t *testing.T) {
	var person plus.Person
	s, _ := ara.Connect("http://localhost:8529", "", "", false)
	user := NewUserModel(s)
	u2 := NewUserModel(s)
	person.Language = "en"
	u2.SetUserName("Ximo")
	user.SetUserName("Miguel")
	user.save()
	u2.save()
	fmt.Printf("\n>>>\nDocment :%+v \n>>>>>\n ", u2.Doc.Age)
}

func TestGetByGoogleId(t *testing.T) {
	s, _ := ara.Connect("http://localhost:8529", "", "", false)
	Z = s
	user, _ := GetByGoogleId("115601102326911748945")
	if user.Doc.Key == "" {
		t.Error("Cant get user by google id")
	}
	fmt.Printf("\n >>> Google >> \n %+v\n", user)
}
