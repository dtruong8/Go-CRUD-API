package main

import (
	"customers/db"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "psqldev"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	store, err := db.NewStore(dsn)
	if err != nil {
		fmt.Printf(err.Error())
	}
	/*db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Connected to Postgres Database!")

	*/
}
