package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//type Users struct {
//	Users []User `json:"users"`
//}

type User struct {
	UserId      string                   `json:"userId"`
	Name        string                   `json:"name"`
	Permissions []map[string]interface{} `json:"permissions"`
}

type Service struct {
	Name        string                 `json:"name"`
	Permissions map[string]interface{} `json:"permissions"`
}

type Permissions struct {
	p map[string]interface{}
}

type BlockbustersPermissions struct {
	Film     bool
	Schedule bool
	Edit     bool
}

type CommercialsPermissions struct {
	License  string
	Monetize int
	Direct   bool
}

type ShortsPermissions struct {
	Cast string
	Fund int
	Act  bool
}

func jsonUnmarshal() {

	ds, err := os.Open("anothertry.json")
	if err != nil {
		return
	}
	defer ds.Close()
	datastore, _ := ioutil.ReadAll(ds)
	fmt.Println(string(datastore))

	var users []User
	err = json.Unmarshal(datastore, &users)
	if err != nil {
		fmt.Println("users unmarshal problem")
		return
	}

}
