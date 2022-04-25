package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)



func truncate(db migrations.DB) {
	fmt.Println("truncate tables...")
	_, e := db.Exec("TRUNCATE CLIENT_VIEW_USERS RESTART IDENTITY CASCADE;")
	panicIfErr(e)
}

func panicIfErr(err error) {
	if err != nil {
		exitf(err.Error())
	}
}
