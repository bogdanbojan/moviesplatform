package db

import (
	_ "embed"
	"encoding/json"
)

//go:embed datastore.json
var datastore []byte

var UsersCollection Users

type Users map[string]User

type User struct {
	Name        string
	Permissions Permissions
}

type Permissions map[string]interface{}

func JsonUnmarshal() {
	err := json.Unmarshal(datastore, &UsersCollection)
	if err != nil {
		return
	}
}
