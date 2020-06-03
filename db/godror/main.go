package main

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

func main() {
	db, err := sql.Open("godror", "FCR_READ_ONLY/readonly@10.55.168.201:1521/FCR1")
	if err != nil {
		panic(err)
	}
	defer db.Close()


	rows, err := db.Query("SELECT addr, HOUSE_ID FROM mv_houses_adreses where HOUSE_ID = 2361")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	type R struct {
		addr    string
		houseId int
	}

	for rows.Next() {
		r := new(R)
		if err := rows.Scan(&r.addr, &r.houseId); err != nil {
			panic(err)
		}
		fmt.Println(*r)
	}

}




