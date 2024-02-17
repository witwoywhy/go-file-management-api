package integrationtest

import (
	"fmt"
	"net/http"
	"testing"
)

func list(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v1/list/%s", baseUrl, userID), nil)
	must(t, err)

	res, err := client.Do(req)
	must(t, err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Error("get list failed")
	}
}
