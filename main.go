// postgresql
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	uri := "postgresql://zollsoft_readonly:tomedo@tomedo/tomedo"
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}

	query := "select kuerzel,vorname,nachname from nutzer where visible = true and kuerzel not in ('test','zollsoft','admin') order by kuerzel"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var kuerzel, vorname, nachname string
	for rows.Next() {
		rows.Scan(&kuerzel, &vorname, &nachname)
		fmt.Println(kuerzel, vorname, nachname)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
