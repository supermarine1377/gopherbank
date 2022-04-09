package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"supermarine1377/domain"
	"supermarine1377/interface/controller"
	"supermarine1377/interface/db"
)

var userController *controller.UserController

func Route(sql db.SqlHandler) {
	userController = controller.NewUserController(sql)
	log.Println("routing...")

	http.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "pong")
	})
	http.HandleFunc("/users", userHandlerFunc)
	http.HandleFunc("/transactions", transactionHandlerFunc)
	http.ListenAndServe(":8080", nil)
	log.Printf("listening...")
}

// func init() {
// 	var (
// 		sql = NewSqlHandler()
// 		db  = sql.Comn
// 	)
// 	if err := db.Ping(); err != nil {
// 		panic(err)
// 	}
// 	userController = controller.NewUserController(sql)
// }

func userHandlerFunc(rw http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "POST" {
		var user domain.User
		dec := json.NewDecoder(r.Body)
		dec.Decode(&user)
		if err := userController.Add(user); err != nil {
			// todo エラーハンドリング
			log.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
		}
		rw.WriteHeader(http.StatusCreated)
	} else if method == "PUT" {
		rw.WriteHeader(http.StatusOK)
	} else if method == "GET" {
		users, err := userController.FindAll()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(rw)
		enc.Encode(&users)
		// rw.WriteHeader(http.StatusOK)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func transactionHandlerFunc(rw http.ResponseWriter, r *http.Request) {

}
