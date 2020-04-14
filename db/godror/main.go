package main

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

func main() {
	db, err := sql.Open("godror", "FCR/Qpwo1029@10.55.168.201:1521/FCR1")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT addr FROM mv_houses_adreses where HOUSE_ID = 2361")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			panic(err)
		}
		fmt.Println(s)
	}

}
