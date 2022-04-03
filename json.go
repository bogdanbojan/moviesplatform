package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	UserId      string      `json:"userId"`
	Name        string      `json:"name"`
	Permissions Permissions `json:"permissions"`
}

type Permissions struct {
	Create  interface{} `json:"create"`
	Modify  interface{} `json:"modify"`
	Comment interface{} `json:"comment"`
	Rate    interface{} `json:"rate"`
	None    interface{} `json:"none"`
}

func jsonUnmarshal() {

	ds, err := os.Open("datastore.json")
	if err != nil {
		return
	}
	defer ds.Close()

	datastore, _ := ioutil.ReadAll(ds)
	fmt.Println(string(datastore))
	var users Users
	err = json.Unmarshal(datastore, &users)
	if err != nil {
		fmt.Println("oops")
		return
	}

	for _, u := range users.Users {
		fmt.Println(u)
	}
}
