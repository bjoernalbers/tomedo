// tomedo displays all tomedo users (only real people w/o special users)
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bjoernalbers/tomedo"
)

func init() {
	log.SetPrefix(fmt.Sprintf("%s: ", filepath.Base(os.Args[0])))
	log.SetFlags(0)
}

func main() {
	users, err := tomedo.Users()
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		fmt.Println(u)
	}
}
