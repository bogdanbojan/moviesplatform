package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserPermissions(t *testing.T) {
	t.Parallel()
	StorageStub.InitStorage("datastore_test.json")
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
	t.Parallel()
	StorageStub.InitStorage("datastore_test.json")
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
}
var PermissionsURLResponsesCases = struct {
	userPermissionsResponses    []int
	servicePermissionsResponses []int
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
}
