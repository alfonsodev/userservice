package userservice

import (
	"code.google.com/p/google-api-go-client/plus/v1"
	"errors"
	//	"fmt"
	ara "github.com/diegogub/aranGO"
)

var Z *ara.Session

type UserDocument struct {
	ara.Document // Must include arango Document in every struct you want to save id, key, rev after saving it
	Username     string
	Age          int
	Likes        []string
	Person       plus.Person
}

type UserModel struct {
	s   *ara.Session
	Doc UserDocument
}

func NewUserModel(s *ara.Session) UserModel {
	var um UserModel
	um.s = s
	return um
}

func (um *UserModel) SetUserName(name string) {
	um.Doc.Username = name
}

func (um *UserModel) save() {
	um.s.DB(USER_DB_NAME).Col(USER_COL_NAME).Save(&um.Doc)
}

func GetByGoogleId(googleId string) (userModel UserModel, err error) {
	if googleId == "" {
		return userModel, errors.New("GooglPlusId can't be empty")
	}
	aql := "FOR c IN users FILTER c.Person.id == \"" + googleId + "\" RETURN c"
	//	fmt.Printf("\nquery:%s\n", aql)
	q := ara.NewQuery(aql)
	c, err := Z.DB(USER_DB_NAME).Execute(q)
	if err != nil {
		panic(err)
	}
	c.FetchOne(&userModel.Doc)

	return userModel, err
}
