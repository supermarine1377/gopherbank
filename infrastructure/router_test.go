package infrastructure_test

import (
	"net/http"
	"net/http/httptest"
	"supermarine1377/domain"
	"supermarine1377/infrastructure"

	"supermarine1377/infrastructure/mock"

	"testing"

	"github.com/golang/mock/gomock"
)

func TestPingHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		statusCode int
	}{
		{
			name:       "GET",
			method:     "GET",
			statusCode: http.StatusOK,
		},
		{
			name:       "POST",
			method:     "POST",
			statusCode: http.StatusMethodNotAllowed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				ctrl = gomock.NewController(t)
				uc   = mock.NewMockUserController(ctrl)
				ro   = infrastructure.NewRouter(uc)
			)
			var (
				w   = httptest.NewRecorder()
				req = httptest.NewRequest(tt.method, "/", nil)
			)
			ro.PingHandler(w, req)
			result := w.Result()
			defer result.Body.Close()
			if result.StatusCode != tt.statusCode {
				t.Errorf("error: unexpectd status code %d, expected %d", result.StatusCode, tt.statusCode)
			}
		})
	}
}

func TestUserHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		param      interface{}
		prepare    func(*mock.MockUserController, domain.User)
		res        domain.User
		statusCode int
	}{
		{
			name:   "GET",
			method: "GET",
			param:  1,
			prepare: func(c *mock.MockUserController, u domain.User) {
				c.EXPECT().FindById(1).Return(&u, nil)
			},
			statusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				ctrl = gomock.NewController(t)
				uc   = mock.NewMockUserController(ctrl)
				ro   = infrastructure.NewRouter(uc)
			)
			rec := httptest.NewRecorder()
			var req *http.Request
			if tt.method == "GET" {
				tt.prepare(uc, tt.res)
				req = httptest.NewRequest(tt.method, "/user/1", nil)
			}
			ro.UserHandler(rec, req)

			result := rec.Result()
			if result.StatusCode != tt.statusCode {
				t.Errorf("unexpected status code %d, expected %d", result.StatusCode, tt.statusCode)
			}
		})
	}
}
