package web

import (
	"reflect"
	"testing"
)

func TestExtractData(t *testing.T) {
	for _, edc := range extractDataCases {
		t.Run(edc.name, func(t *testing.T) {
			u, s := extractData(edc.url)
			assertEqualString(t, u, edc.wantUser)
			assertEqualString(t, s, edc.wantService)
		})
	}
}

var extractDataCases = []struct {
	name        string
	url         string
	wantUser    string
	wantService string
}{
	{
		name:        "user and service",
		url:         "v1/user/user2411/blockbusters",
		wantUser:    "user2411",
		wantService: "blockbusters",
	},
	{
		name:        "multiple slashes at the end",
		url:         "v1/user/user2411/blockbusters///",
		wantUser:    "user2411",
		wantService: "blockbusters",
	},
	{
		name:        "another user and service",
		url:         "v1/user/user211/commercials",
		wantUser:    "user211",
		wantService: "commercials",
	},
}

func TestConstructServicePermissionData(t *testing.T) {
	StorageStub.InitStorage(TestStorage)
	for _, spdc := range constructServicePermissionDataCases {
		t.Run(spdc.name, func(t *testing.T) {
			got := AppStub.constructServicePermissionData(spdc.userId, spdc.service)
			assertEqualPermission(t, got, spdc.want)
		})
	}
}

func assertEqualPermission(t testing.TB, got, want map[string]interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}

var constructServicePermissionDataCases = []struct {
	name    string
	userId  string
	service string
	want    map[string]interface{}
}{
	{
		name:    "valid user",
		userId:  "user2143",
		service: "blockbusters",
		want:    map[string]interface{}{"blockbusters.cinematographer.changeLens": true, "blockbusters.cinematographer.shoot": false, "blockbusters.cinematographer.changeCamera": "allow"},
	},
	{
		name:    "valid user without permissions in that particular service",
		userId:  "user4323",
		service: "commercials",
		want:    map[string]interface{}{},
	},
	{
		name:    "user that does not exist",
		userId:  "user2d23143",
		service: "blockbusters",
		want:    map[string]interface{}{},
	},
	{
		name:    "random strings",
		userId:  "jafdssdah12",
		service: "sdfgjklad",
		want:    map[string]interface{}{},
	},
}

func assertEqualString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
