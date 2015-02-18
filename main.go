package userservice

import (
	"errors"
	"fmt"
	ara "github.com/diegogub/aranGO"
)

func main() {
	fmt.Printf("\nHello user service...\n>>>>>\n")
}

func NewUserModel() UserModel {
	var um UserModel
	return um
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
		return userModel, err
	}
	c.FetchOne(&userModel.Doc)

	return userModel, err
}
