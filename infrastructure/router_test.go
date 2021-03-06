package infrastructure_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
			res: domain.User{
				Name: "gopher",
			},
			statusCode: http.StatusOK,
		},
		{
			name:   "GET (Not found)",
			method: "GET",
			prepare: func(muc *mock.MockUserController, u *domain.User) {
				muc.EXPECT().FindById(1).Return(u, domain.ErrUserNotFound(1))
			},
			statusCode: http.StatusNotFound,
		},
		{
			name:   "POST (sucessful)",
			method: "POST",
			req: `{
				"name": "gopher",
				"balance": 10000
			}`,
			prepare: func(muc *mock.MockUserController, u *domain.User) {
				muc.EXPECT().Add(*u).Return(nil)
			},
			statusCode: http.StatusCreated,
		},
		{
			name:   "POST (invalid balance)",
			method: "POST",
			req: `{
				"name": "gopher",
				"balance": -10000
			}`,
			prepare: func(muc *mock.MockUserController, u *domain.User) {
				muc.EXPECT().Add(*u).Return(domain.ErrInvalidUserCreateReq)
			},
			statusCode: http.StatusBadRequest,
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
				rec  = httptest.NewRecorder()
				req  *http.Request
			)

			switch tt.method {
			case "GET":
				tt.prepare(uc, &tt.res)
				req = httptest.NewRequest("GET", "/user/1", nil)
			case "POST":
				reqBody := strings.NewReader(tt.req)
				dec := json.NewDecoder(reqBody)
				var user domain.User
				dec.Decode(&user)
				tt.prepare(uc, &user)
				req = httptest.NewRequest("POST", "/user", bytes.NewBufferString(tt.req))
				fmt.Println(req.Body)
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

func TestUsersAllHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		prepare    func(*mock.MockUserController, []domain.User)
		res        []domain.User
		statusCode int
	}{
		{
			name:   "get(successful)",
			method: "GET",
			prepare: func(muc *mock.MockUserController, u []domain.User) {
				muc.EXPECT().FindAll().Return(u, nil)
			},
			res: []domain.User{
				{
					ID:      1,
					Name:    "gopher1",
					Balance: 100,
				},
				{
					ID:      2,
					Name:    "gopher2",
					Balance: 100,
				},
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "post",
			method:     "POST",
			statusCode: http.StatusNotImplemented,
		},
		{
			name:       "put",
			method:     "PUT",
			statusCode: http.StatusNotImplemented,
		},
	}
	for _, tt := range tests {
		var (
			ctrl = gomock.NewController(t)
			muc  = mock.NewMockUserController(ctrl)
			ro   = infrastructure.NewRouter(muc)
		)
		rec := httptest.NewRecorder()
		var req *http.Request

		switch tt.method {
		case "GET":
			req = httptest.NewRequest("GET", "/users", nil)
			tt.prepare(muc, tt.res)
		default:
			req = httptest.NewRequest(tt.method, "/users", nil)
		}

		ro.UsersAllHandler(rec, req)
		t.Run(tt.name, func(t *testing.T) {
			result := rec.Result()
			if result.StatusCode != tt.statusCode {
				t.Errorf("error: unexpected stauts code %d, expected %d", result.StatusCode, tt.statusCode)
			}
		})
	}
}
