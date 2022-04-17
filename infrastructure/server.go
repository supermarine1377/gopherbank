package infrastructure

import (
	"database/sql"
	"log"
	"net/http"
	"supermarine1377/interface/controller"
)

func Run() {
	sh, db := connectDB()
	defer db.Close()
	var (
		uc = controller.NewUserController(sh)
		r  = NewRouter(uc)
	)
	useRouter(r)
	http.ListenAndServe(":8080", nil)
	log.Printf("listening...")
}

func connectDB() (*SqlHandler, *sql.DB) {
	var (
		sql = NewSqlHandler()
		db  = sql.Comn
	)
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return sql, db
}

func useRouter(r *Router) {
	http.HandleFunc("/ping", r.PingHandler)
	http.HandleFunc("/users", r.UsersAllHandler)
	http.HandleFunc("/user/", r.UserHandler)
}
