package userservice

import (
	ara "github.com/diegogub/aranGO"
)

var Z *ara.Session

func init() {
	//s, err := ara.Connect("http://localhost:8529", "", "", false)
	s, err := ara.Connect(ARANGO_SERVER, ARANGO_USER, ARANGO_PASS, false)
	if err != nil {
		panic("Can't connect to ")
	}
	Z = s
}
