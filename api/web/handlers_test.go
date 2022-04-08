package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserPermissions(t *testing.T) {
	StorageStub.InitStorage(TestStorage)
	for i, urc := range PermissionsURLRequestsCases.userPermissions {
		rr := httptest.NewRecorder()
		r, err := http.NewRequest(PermissionsURLRequestsCases.methodRequests[i], urc, nil)
		if err != nil {
			t.Fatal(err)
		}
		AppStub.Permissions(rr, r)
		rs := rr.Result()
		assertStatusCode(t, rs.StatusCode, PermissionsURLResponsesCases.userPermissionsResponses[i])
		err = rs.Body.Close()
		if err != nil {
			return
		}
	}
}

func TestServicePermissions(t *testing.T) {
	StorageStub.InitStorage(TestStorage)
	for i, urc := range PermissionsURLRequestsCases.servicePermissions {
		rr := httptest.NewRecorder()
		r, err := http.NewRequest(PermissionsURLRequestsCases.methodRequests[i], urc, nil)
		if err != nil {
			t.Fatal(err)
		}
		AppStub.Permissions(rr, r)
		rs := rr.Result()
		assertStatusCode(t, rs.StatusCode, PermissionsURLResponsesCases.servicePermissionsResponses[i])
		err = rs.Body.Close()
		if err != nil {
			return
		}
	}
}

func TestUsersPermissions(t *testing.T) {
	StorageStub.InitStorage(TestStorage)
	for i, urc := range PermissionsURLRequestsCases.sfp {
		rr := httptest.NewRecorder()
		r, err := http.NewRequest(PermissionsURLRequestsCases.methodRequests[i], urc, nil)
		if err != nil {
			t.Fatal(err)
		}
		AppStub.Users(rr, r)
		rs := rr.Result()
		assertStatusCode(t, rs.StatusCode, PermissionsURLResponsesCases.sfpResponses[i])
		err = rs.Body.Close()
		if err != nil {
			return
		}
	}
}

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}

var PermissionsURLRequestsCases = struct {
	methodRequests     []string
	userPermissions    []string
	servicePermissions []string
	sfp                []string
}{
	methodRequests: []string{
		"GET",
		"POST",
		"GET",
		"GET",
		"GET",
		"GET",
		"GET",
	},
	userPermissions: []string{
		"/v1/user/user2143",
		"/v1/user/user2143",
		"/v1/user/user214312",
		"/v2/user/user2143",
		"/fdwaj231",
		"/v1/user/user2143/dasf/12",
		"/v1/user/uSer2143",
	},
	servicePermissions: []string{
		"/v1/user/user2143/blockbusters",
		"/v1/user/user2143/blockbusters",
		"/v1/user/user214312/blockbusters",
		"/v2/user/user2143/commercials",
		"/fdwaj231",
		"/v1/user/user2143/dasf/12",
		"/v1/user/user2143/commerCials",
	},
	sfp: []string{
		"/v1/service/blockbusters/cinematographer/shoot",
		"/v1/service/blockbusters/commercials/shoot",
		"/v1/service/blockBusters/cinematographer/shoot",
		"/v2/user/user2143/commercials",
		"/fdwaj231",
		"/v1/service/blockbusters/cinematographersds/",
		"/v1/service/commercials/manageR/shoot",
	},
}
var PermissionsURLResponsesCases = struct {
	userPermissionsResponses    []int
	servicePermissionsResponses []int
	sfpResponses                []int
}{
	userPermissionsResponses: []int{
		http.StatusOK,
		http.StatusMethodNotAllowed,
		http.StatusBadRequest,
		http.StatusBadRequest,
		http.StatusBadRequest,
		http.StatusBadRequest,
		http.StatusBadRequest,
	},
	servicePermissionsResponses: []int{
		http.StatusOK,
		http.StatusMethodNotAllowed,
		http.StatusBadRequest,
		http.StatusBadRequest,
		http.StatusBadRequest,
		http.StatusBadRequest,
		http.StatusBadRequest,
	},
	sfpResponses: []int{
		http.StatusOK,
		http.StatusMethodNotAllowed,
		http.StatusBadRequest,
		http.StatusBadRequest,
		http.StatusBadRequest,
		http.StatusBadRequest,
		http.StatusBadRequest,
	},
}
