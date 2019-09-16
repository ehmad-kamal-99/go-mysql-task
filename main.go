package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Docker Practise")

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3308)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Connection Successfull!")

	f, _ := os.Open("user_data.csv")

	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		sqlstatement := `
		INSERT INTO user (firstname, lastname, age, bloodgroup)
		VALUES (?, ?, ?, ?)
		`

		_, err = db.Exec(sqlstatement, record[0], record[1], record[2], record[3])

		if err != nil {
			panic(err.Error())
		}
	}
}
