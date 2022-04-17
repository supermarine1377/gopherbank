package infrastructure_test

import (
	"net/http"
	"net/http/httptest"
	"supermarine1377/infrastructure"

	"supermarine1377/interface/controller"
	mock_db "supermarine1377/interface/db/mock"

	"testing"

	"github.com/golang/mock/gomock"
)

func TestPingHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		prepare    func(ctrl *gomock.Controller) *controller.UserController
		statusCode int
	}{
		{
			name:   "GET",
			method: "GET",
			prepare: func(ctrl *gomock.Controller) *controller.UserController {
				msh := mock_db.NewMockSqlHandler(ctrl)
				return controller.NewUserController(msh)
			},
			statusCode: http.StatusOK,
		},
		{
			name:   "POST",
			method: "POST",
			prepare: func(ctrl *gomock.Controller) *controller.UserController {
				msh := mock_db.NewMockSqlHandler(ctrl)
				return controller.NewUserController(msh)
			},
			statusCode: http.StatusMethodNotAllowed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				ctrl = gomock.NewController(t)
				uc   = tt.prepare(ctrl)
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
	type fields struct {
		userController controller.UserController
	}
	type args struct {
		rw http.ResponseWriter
		r  *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
