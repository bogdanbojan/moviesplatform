package web

import (
	"github.com/bogdanbojan/moviesplatform/db"
	"testing"
)

var StorageStub = &db.Storage{
	ServicesStructure: db.InitServiceStructure(),
}
var AppStub = Application{
	ServiceLogger: nil,
	DataPuller:    StorageStub,
}

func TestCheckServiceRoute(t *testing.T) {
	for _, sc := range checkServiceRouteCases {
		t.Run(sc.name, func(t *testing.T) {
			got := AppStub.checkServiceRoute(sc.serviceRoutes)
			want := sc.want
			assertEqual(t, got, want)
		})
	}
}

func assertEqual(t testing.TB, got bool, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

var checkServiceRouteCases = []struct {
	name          string
	serviceRoutes []string
	want          bool
}{
	{
		name:          "valid service route",
		serviceRoutes: []string{"blockbusters", "cinematographer", "shoot"},
		want:          true,
	},
	{
		name:          "feature from another service",
		serviceRoutes: []string{"blockbusters", "manager", "shoot"},
		want:          false,
	},
	{
		name:          "misspelled permission",
		serviceRoutes: []string{"blockbusters", "cinematographer", "shoOt"},
		want:          false,
	},
	{
		name:          "non existent permission",
		serviceRoutes: []string{"blockbusters", "cinematographer", "articulate"},
		want:          false,
	},
	{
		name:          "random strings",
		serviceRoutes: []string{"fadsjk213", "dfas12", "fdsda21ash21"},
		want:          false,
	},
}

func TestCheckService(t *testing.T) {
	for _, sc := range checkServiceCases {
		t.Run(sc.name, func(t *testing.T) {
			got := checkService(sc.service)
			want := sc.want
			assertEqual(t, got, want)
		})
	}
}

var checkServiceCases = []struct {
	name    string
	service string
	want    bool
}{
	{
		name:    "valid service",
		service: "blockbusters",
		want:    true,
	},
	{
		name:    "input is a feature",
		service: "director",
		want:    false,
	},
	{
		name:    "misspelled service",
		service: "comercials",
		want:    false,
	},
	{
		name:    "random string",
		service: "dsaj2k212",
		want:    false,
	},
}

func TestValidateUsersURL(t *testing.T) {
	for _, sc := range validateUsersURLCases {
		t.Run(sc.name, func(t *testing.T) {
			got := AppStub.validateUsersURL(sc.url)
			want := sc.want
			assertEqual(t, got, want)
		})
	}
}

var validateUsersURLCases = []struct {
	name string
	url  string
	want bool
}{
	{
		name: "valid url",
		url:  "/v1/service/blockbusters/cinematographer/shoot",
		want: true,
	},
	{
		name: "feature from another url",
		url:  "/v1/service/blockbusters/director/shoot",
		want: false,
	},
	{
		name: "misspelled permission",
		url:  "/v1/service.blockbusters/director/shOot",
		want: false,
	},
	{
		name: "random string",
		url:  "/v1/23dsfae/sads/dafdsa21/shOotsa/",
		want: false,
	},
}

func TestValidatePermissionsURL(t *testing.T) {
	for _, sc := range validatePermissionsURLCases {
		t.Run(sc.name, func(t *testing.T) {
			got := AppStub.validatePermissionsURL(sc.url)
			want := sc.want
			assertEqual(t, got, want)
		})
	}
}

var validatePermissionsURLCases = []struct {
	name string
	url  string
	want bool
}{
	{
		name: "random string",
		url:  "/v1/23dsfae/sads/shOotsa/",
		want: false,
	},
	{
		name: "misspelled user",
		url:  "/v1/user/uSe2143/",
		want: false,
	},
	{
		name: "nonexistent user",
		url:  "/v1/user/user2110/",
		want: false,
	},
}
