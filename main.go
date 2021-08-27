package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:13306)/mysql")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select Host, User from user")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for i := 0; i < 100; i++ {
		var host string
		var user string
		if err := db.QueryRow("select Host, user from user limit 1").Scan(&host, &user); err != nil {
			panic(err)
		}
		fmt.Printf("%05d: (Host, User) ... (%s, %s)\n", i, host, user)

		time.Sleep(time.Second * 1)
	}
}
