package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table cars...")
		_, err := db.Exec(`CREATE TABLE cars (
		id SERIAL PRIMARY KEY,
		make TEXT NOT NULL,
		model TEXT NOT NULL,
		year INT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table cars...")
		_, err := db.Exec(`DROP TABLE cars`)
		return err
	})
}
