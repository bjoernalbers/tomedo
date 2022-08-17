// postgresql
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	/*
			query := "select kuerzel,vorname,nachname from nutzer where visible = true and kuerzel not in ('test','zollsoft','admin') order by kuerzel"
	    uri := "postgresql://zollsoft_readonly:tomedo@tomedo/tomedo"
			cmd := exec.Command("psql", "--command", query, "--tuples-only", "--csv", uri)
			output, err := cmd.Output()
			fmt.Print(string(output))
			if err != nil {
				fmt.Println(err)
			}
	*/
	uri := "postgresql://zollsoft_readonly:tomedo@tomedo/tomedo"
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db.Ping())
}
