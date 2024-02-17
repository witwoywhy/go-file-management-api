package integrationtest

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func upload(t *testing.T) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	file, err := os.Open(fileName)
	must(t, err)
	defer file.Close()

	part, err := writer.CreateFormFile("file", filepath.Base(fileName))
	must(t, err)

	_, err = io.Copy(part, file)
	must(t, err)
	writer.Close()

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v1/upload/%s", baseUrl, userID), payload)
	must(t, err)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	must(t, err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Error("upload file failed")
	}
}
