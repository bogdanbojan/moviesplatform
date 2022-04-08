package db

import (
	_ "embed"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

const (
	Embed = ""
	File  = "datastore.json"
)

//go:embed datastore.json
var datastore []byte

// UsersCollection represents our database of users.
var UsersCollection Users

type Users map[string]User

type User struct {
	Name        string
	Permissions Permissions
}

type Permissions map[string]interface{}

func (s *Storage) JSONUnmarshalEmbed() {
	err := json.Unmarshal(datastore, &UsersCollection)
	if err != nil {
		return
	}
}

func (s *Storage) JSONUnmarshalFile(fileName string) {
	_, b, _, _ := runtime.Caller(0)
	currDir := path.Join(path.Dir(b))
	abs, err := filepath.Abs(currDir + `\` + fileName)
	if err != nil {
		return
	}
	f, err := os.Open(abs)
	if err != nil {
		return
	}
	defer f.Close()
	ds, _ := ioutil.ReadAll(f)

	err = json.Unmarshal(ds, &UsersCollection)
	if err != nil {
		return
	}

}
