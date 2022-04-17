//go:generate mockgen -source=$GOFILE -package=mock -destination=./mock/mock_$GOFILE
package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"supermarine1377/domain"
	"supermarine1377/interface/controller"
)

type Router struct {
	userController controller.UserController
}

type IRouter interface {
	PingHandler(w io.Writer, r io.Reader)
}

func NewRouter(uc *controller.UserController) *Router {
	return &Router{userController: *uc}
}

func (ro *Router) PingHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(rw, "pong")
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (ro *Router) UserHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		enc := json.NewEncoder(rw)
		prx := strings.TrimPrefix(r.URL.Path, "/user/")
		id, err := strconv.Atoi(prx)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		user, err := ro.userController.FindById(id)
		if err != nil {
			if err.Error() == domain.UserNotFoundErr(id).Error() {
				rw.WriteHeader(http.StatusNotFound)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		enc.Encode(&user)
	case "POST":
		var user domain.User
		dec := json.NewDecoder(r.Body)
		dec.Decode(&user)
		if err := ro.userController.Add(user); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusCreated)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (ro *Router) UsersAllHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		enc := json.NewEncoder(rw)
		users, err := ro.userController.FindAll()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		enc.Encode(&users)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
