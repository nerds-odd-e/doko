package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table users...")
		_, err := db.Exec(`
			CREATE TABLE users(
			    id serial PRIMARY KEY,
			    username VARCHAR (500) UNIQUE NOT NULL,
			    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			    updated_by integer NULL
			);
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table users...")
		_, err := db.Exec(`
			DROP TABLE users;
		`)
		return err
	})
}
