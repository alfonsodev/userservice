package userservice

import (
	"os"
	"strings"
)

const (
	USER_DB_NAME  = "user-managment"
	USER_COL_NAME = "users"
)

var (
	clientID        = os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret    = os.Getenv("GOOGLE_CLIENT_SECRET")
	applicationName = ""
	ARANGO_SERVER   = strings.Replace(os.Getenv("ARANGO_PORT"), "tcp", "http", -1)
	ARANGO_USER     = os.Getenv("ARANGO_USER")
	ARANGO_PASS     = os.Getenv("ARANGO_PASS")
)
