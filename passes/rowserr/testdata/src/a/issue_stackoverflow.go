package a

import (
	"database/sql"
	"fmt"
)

func issue_stackoverflow() {
	db, err := sql.Open("postgres", "postgres://localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	getRows(rows)
}

func getRows(rows *sql.Rows) {
	for rows.Next() {
		fmt.Println("new rows")
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
}
