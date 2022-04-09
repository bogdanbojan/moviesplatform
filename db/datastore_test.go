package db

import (
	"testing"
)

const TestStorage = "datastore_test.json"

var StorageStub = &Storage{
	ServicesStructure: InitServiceStructure(),
}

func TestGetUser(t *testing.T) {
	StorageStub.InitStorage(TestStorage)
	for _, uc := range getUserCases {
		t.Run(uc.name, func(t *testing.T) {
			u, b := StorageStub.GetUser(uc.user)
			assertEqualString(t, u.Name, uc.wantUser)
			assertEqualBool(t, b, uc.wantBool)
		})
	}
}

var getUserCases = []struct {
	name     string
	user     string
	wantUser string
	wantBool bool
}{
	{
		name:     "valid user",
		user:     "user2143",
		wantUser: "Roger Deakins",
		wantBool: true,
	},
	{
		name:     "user does not exist",
		user:     "user21432sf1",
		wantUser: "",
		wantBool: false,
	},
	{
		name:     "random string",
		user:     "fdgsoa123",
		wantUser: "",
		wantBool: false,
	},
	{
		name:     "misspelled user",
		user:     "useR2143",
		wantUser: "",
		wantBool: false,
	},
}

func TestValidateServicePermission(t *testing.T) {
	for _, spc := range validateServicePermissionCases {
		t.Run(spc.name, func(t *testing.T) {
			got := StorageStub.ValidateServiceFeaturePermission(spc.sfp)
			assertEqualBool(t, got, spc.want)
		})
	}
}

var validateServicePermissionCases = []struct {
	name string
	sfp  []string
	want bool
}{
	{
		name: "valid service",
		sfp:  []string{"blockbusters", "cinematographer", "shoot"},
		want: true,
	},
	{
		name: "feature from another service",
		sfp:  []string{"blockbusters", "manager", "shoot"},
		want: false,
	},
	{
		name: "misspelled permission",
		sfp:  []string{"blockbusters", "cinematographer", "shoOt"},
		want: false,
	},
	{
		name: "random strings",
		sfp:  []string{"fdsads", "fds", "ds321"},
		want: false,
	},
}

func assertEqualString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}

func assertEqualBool(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}

}
