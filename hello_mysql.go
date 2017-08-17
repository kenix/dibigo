package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:dba@tcp(:3306)/test")
	check(err)
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test.hello(world varchar(50))")
	check(err)

	res, err := db.Exec("INSERT INTO test.hello(world) values('hello world!')")
	check(err)

	rowCount, err := res.RowsAffected()
	check(err)
	log.Printf("inserted %d row(s)", rowCount)

	rows, err := db.Query("SELECT * FROM test.hello")
	defer rows.Close()
	check(err)
	for rows.Next() {
		var s string
		err := rows.Scan(&s)
		check(err)
		log.Printf("found row containing %q", s)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
