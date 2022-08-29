// tomedo allows access to tomedo's database
package tomedo

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	Name      string
	FirstName string
	LastName  string
}

func (u User) String() string {
	return fmt.Sprintf("%s %s (%s)", u.FirstName, u.LastName, u.Name)
}

// tomedoUsers returns all users found in tomedo's database
func Users() ([]User, error) {
	var users []User

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
		var u User
		rows.Scan(&u.Name, &u.FirstName, &u.LastName)
		users = append(users, u)
	}
	return users, rows.Err()
}
