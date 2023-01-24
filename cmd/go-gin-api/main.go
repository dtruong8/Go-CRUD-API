package main

import (
	"customers/db"
	"fmt"
	"net/http"
)

const webPort = "80"

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

	http.HandleFunc("/", store.CustomerStore.Customers())
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf(err.Error())
	}

}
