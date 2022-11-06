package main

import (
	// "example/hello/config"
	"example/hello/config"
	"fmt"
	"log"
	"net/http"
)

// /home/didik/Workspace/go_path/src/github.com/kodingin/api-mysql
// github.com/kodingin/api-mysql/Config
// GOPATH/src/github.com/iniaku/api-mysql

func main() {
	db, e := config.MsSql()

	if e != nil {
		panic(e)
	}
	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("database terhubung")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
