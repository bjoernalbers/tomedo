// postgresql
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	query := "select kuerzel,vorname,nachname from nutzer where visible = true and kuerzel not in ('test','zollsoft','admin') order by kuerzel"
	uri := "postgresql://zollsoft_readonly:tomedo@tomedo/tomedo"
	cmd := exec.Command("psql", "--command", query, "--tuples-only", "--csv", uri)
	output, err := cmd.Output()
	fmt.Print(string(output))
	if err != nil {
		fmt.Println(err)
	}
}
