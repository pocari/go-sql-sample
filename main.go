package main

import (
	"context"
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

	// context.WithCancelとかを使ってキャンセルできるようになってるらしい
	// 1秒後に実行中のsqlをキャンセルする、みたいなことができる
	//
	// ctx, cancel := context.WithCancel(context.Background())
	// go func() {
	// 	time.Sleep(time.Second * 1)
	// 	cancel()
	// }

	ctx := context.Background()
	for i := 0; i < 100; i++ {
		var host string
		var user string
		if err := db.QueryRowContext(ctx, "select Host, user from user limit 1").Scan(&host, &user); err != nil {
			panic(err)
		}
		fmt.Printf("%05d: (Host, User) ... (%s, %s)\n", i, host, user)

		time.Sleep(time.Millisecond * 500)
	}
}
