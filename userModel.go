package userservice

import (
	"code.google.com/p/google-api-go-client/plus/v1"
	ara "github.com/diegogub/aranGO"
)

type UserDocument struct {
	ara.Document // Must include arango Document in every struct you want to save id, key, rev after saving it
	Username     string
	Age          int
	Likes        []string
	Person       plus.Person
}

type UserModel struct {
	Doc UserDocument
}

func (um *UserModel) SetUserName(name string) {
	um.Doc.Username = name
}

func (um *UserModel) save() {
	Z.DB(USER_DB_NAME).Col(USER_COL_NAME).Save(&um.Doc)
}
