package infrastructure

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
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
	fmt.Println("gopherbank started!")
	http.ListenAndServe(":8080", nil)
}

func connectDB() (*SqlHandler, *sql.DB) {
	fmt.Println(os.Getenv("IS_ENV"))
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
