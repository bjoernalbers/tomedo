// tomedo displays all tomedo users (only real people w/o special users)
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

func init() {
	log.SetPrefix(fmt.Sprintf("%s: ", filepath.Base(os.Args[0])))
	log.SetFlags(0)
}

func main() {
	users, err := tomedoUsers()
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		fmt.Println(u)
	}
}

type user struct {
	username string
	forename string
	surname  string
}

func (u user) String() string {
	return fmt.Sprintf("%s %s (%s)", u.forename, u.surname, u.username)
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
