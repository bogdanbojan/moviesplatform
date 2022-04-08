package web

import "testing"

func TestExtractData(t *testing.T) {
	for i, edc := range extractDataCases.cases {
		u, s := extractData(edc)
		assertEqualString(t, u, extractDataResponses.responses[i][0])
		assertEqualString(t, s, extractDataResponses.responses[i][1])
	}
}
func assertEqualString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}

var extractDataCases = struct {
	cases []string
}{
	cases: []string{"v1/user/user2411/blockbusters", "v1/user/user2411/blockbusters///", "v1/user/user211/commercials"},
}

var extractDataResponses = struct {
	responses [][2]string
}{
	responses: [][2]string{{"user2411", "blockbusters"}, {"user2411", "blockbusters"}, {"user211", "commercials"}},
}
