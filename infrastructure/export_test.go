package infrastructure

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandlerFunc(t *testing.T) {
	var testcases = []struct {
		name   string
		method string
		res    struct {
			statusCode int
		}
	}{
		{
			name:   "get",
			method: "GET",
			res: struct{ statusCode int }{
				statusCode: http.StatusMethodNotAllowed,
			},
		},
		{
			name:   "delete",
			method: "DELETE",
			res: struct{ statusCode int }{
				statusCode: http.StatusMethodNotAllowed,
			},
		},
		{
			name:   "post",
			method: "POST",
			res: struct{ statusCode int }{
				statusCode: http.StatusCreated,
			},
		},
		{
			name:   "put",
			method: "PUT",
			res: struct{ statusCode int }{
				statusCode: http.StatusOK,
			},
		},
	}
	for _, tc := range testcases {
		func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(tc.method, "/", nil)
			userHandlerFunc(w, r)
			rw := w.Result()
			defer rw.Body.Close()
			t.Run(tc.name, func(t *testing.T) {
				if rw.StatusCode != tc.res.statusCode {
					t.Errorf("unexpected status code")
				}
			})
		}()
	}
}
