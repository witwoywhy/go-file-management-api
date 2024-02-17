package integrationtest

import (
	"fmt"
	"net/http"
	"testing"
)

func remove(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/file/%s/%s", baseUrl, userID, fileName), nil)
	must(t, err)

	res, err := client.Do(req)
	must(t, err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Error("remove failed failed")
	}
}
