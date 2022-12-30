package commands

import (
	"database/sql"
	"fmt"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

func Sql(ChannelID string, slice []string) {
	if len(slice) == 1 {
		return
	}
	// Create the database handle, confirm driver is present
	db, _ := sql.Open("mysql", "root:password@/mydb")
	defer db.Close()

	// Connect and check the server version
	var version string
	db.QueryRow(slice[1])
	fmt.Println("Connected to: ", version)
}
