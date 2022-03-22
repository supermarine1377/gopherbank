package main

import (
	"fmt"
	"supermarine1377/config"
)

func main() {
	fmt.Println("starting server...")
	db := config.NewDB()
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
