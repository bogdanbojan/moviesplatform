package main

import (
	_ "embed"
	"encoding/json"
)

//go:embed datastore.json
var datastore []byte

type Users map[string]User

type User struct {
	Name        string
	Permissions Permissions
}

type Permissions map[string]interface{}

func jsonUnmarshal() {
	var users Users

	err := json.Unmarshal(datastore, &users)
	if err != nil {
		return
	}

	//
	//for user := range users {
	//	fmt.Println(user)
	//	for permission := range users[user].Permissions {
	//		fmt.Println(users[user].Permissions[permission])
	//	}
	//	fmt.Println()
	//}

}
