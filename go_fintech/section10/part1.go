package section10

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func Main1() {
	DbConnection, _ := sql.Open("sqlite3", "./section10/example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING, 
		age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err = DbConnection.Exec(cmd, "Tori", 26)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, 30, "John")

	cmd = "SELECT * FROM person"
	rows, _ := DbConnection.Query(cmd)

	defer rows.Close()

	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}

	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}
