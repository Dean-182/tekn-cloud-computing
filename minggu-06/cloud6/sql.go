package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Konfigurasi koneksi MySQL
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/coba")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Membuat query
	query := "SELECT * FROM mahasiswa"

	// Menjalankan query
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Membaca hasil query
	for rows.Next() {
		var column1, column2 string
		err := rows.Scan(&column1, &column2)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(column1, column2)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
}
