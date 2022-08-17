// tomedo-users displays all tomedo users (only real people w/o special users)
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type user struct {
	username string
	forename string
	surname  string
}

func main() {
	users, err := tomedoUsers()
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		fmt.Printf("%s %s (%s)\n", u.forename, u.surname, u.username)
	}
}

// tomedoUsers returns all users found in tomedo's database
func tomedoUsers() ([]user, error) {
	var users []user

	uri := "postgresql://zollsoft_readonly:tomedo@tomedo/tomedo"
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return users, err
	}

	query := "select kuerzel,vorname,nachname from nutzer where visible = true and kuerzel not in ('admin','Arzeko','test','zollsoft') order by kuerzel"
	rows, err := db.Query(query)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		rows.Scan(&u.username, &u.forename, &u.surname)
		users = append(users, u)
	}
	return users, rows.Err()
}
