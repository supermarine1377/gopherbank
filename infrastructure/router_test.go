package infrastructure_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
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
		prepare    func(*mock.MockUserController, *domain.User)
		req        string
		res        domain.User
		statusCode int
	}{
		{
			name:   "GET (successful)",
			method: "GET",
			prepare: func(muc *mock.MockUserController, u *domain.User) {
				muc.EXPECT().FindById(1).Return(u, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:   "POST (sucessful)",
			method: "POST",
			req: `{
				"name": "gopher",
				"balance": 10000
			}`,
			prepare: func(muc *mock.MockUserController, u *domain.User) {
				muc.EXPECT().Add(gomock.Any()).Return(nil)
			},
			statusCode: http.StatusCreated,
		},
		{
			name:       "PUT",
			method:     "PUT",
			statusCode: http.StatusNotImplemented,
		},
		{
			name:       "DELETE",
			method:     "DELETE",
			statusCode: http.StatusNotImplemented,
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
			switch tt.method {
			case "GET":
				tt.prepare(uc, &tt.res)
				req = httptest.NewRequest("GET", "/user/1", nil)
			case "POST":
				tt.prepare(uc, nil)
				reqBody := strings.NewReader(tt.req)
				req = httptest.NewRequest("POST", "/user", reqBody)
			default:
				req = httptest.NewRequest(tt.method, "/user", nil)
			}

			ro.UserHandler(rec, req)

			result := rec.Result()
			if result.StatusCode != tt.statusCode {
				t.Errorf("unexpected status code %d, expected %d", result.StatusCode, tt.statusCode)
			}
		})
	}
}
