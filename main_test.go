package main_test

import (
	"bytes"
	"net/http"
	"testing"
)

const (
	baseUrl = "http://localhost:8080"
)

func Test_main(t *testing.T) {
	testUsers(t)
}

// testing for GET /users
func testUsers(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		buffer := bytes.NewBuffer(make([]byte, 128))
		req, err := http.NewRequest("GET", baseUrl+"/users", buffer)
		if err != nil {
			t.Fatal(t)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("error: unexpected status code %d, expected %d", resp.StatusCode, http.StatusOK)
		}
	})
}
