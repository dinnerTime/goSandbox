package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/dinnerTime/goSandbox"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	check(err)
	defer db.Close()

	fmt.Println("writing statement")

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	check(err)

	fmt.Println("Db created, inserting rows...")

	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	check(err)

	defer stmt.Close()
	for i := 0; i < 4; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		check(err)
	}
	tx.Commit()

	fmt.Println("Write complete, selecting.")

	rows, err := db.Query("select id, name from foo")
	check(err)

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		check(err)
		fmt.Println(id, name)
	}
	check(err)

	fmt.Println("Rows scanned.")

	stmt, err = db.Prepare("select name from foo where id = ?")
	check(err)

	fmt.Println("Rows selected, querying row 3.")

	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	check(err)

	fmt.Println("Rows selected.")

	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	check(err)

	fmt.Println("Delete completed, inserting new.")

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	check(err)

	fmt.Println("insert completed.")

	rows, err = db.Query("select id, name from foo")
	check(err)
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	check(err)

	erp := MyFoo //{Bar: "some string"}
	fmt.Sprintf("Value is %s", erp.Bar)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
