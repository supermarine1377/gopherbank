package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"supermarine1377/domain"
)

type Router struct {
	userController UserController
}

func NewRouter(uc UserController) *Router {
	return &Router{userController: uc}
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
			if err.Error() == domain.ErrUserNotFound(id).Error() {
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
		err := ro.userController.Add(user)
		switch err {
		case nil:
			rw.WriteHeader(http.StatusCreated)
		case domain.ErrInvalidUserCreateReq:
			rw.WriteHeader(http.StatusBadRequest)
		default:
			rw.WriteHeader(http.StatusInternalServerError)
		}
	default:
		rw.WriteHeader(http.StatusNotImplemented)
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
		rw.WriteHeader(http.StatusNotImplemented)
	}
}
