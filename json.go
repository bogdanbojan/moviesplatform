package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users map[string]User

type User struct {
	Name        string
	Permissions Permissions
}
type Permissions map[string]interface{}

func jsonUnmarshal() {
	ds, err := os.Open("datastore.json")
	if err != nil {
		return
	}
	defer ds.Close()
	datastore, _ := ioutil.ReadAll(ds)

	var users Users

	err = json.Unmarshal(datastore, &users)
	if err != nil {
		fmt.Println("users unmarshal problem")
		return
	}

	for user := range users {
		fmt.Println(user)
		for permission := range users[user].Permissions {
			fmt.Println(users[user].Permissions[permission])
		}
		fmt.Println()
	}

}
