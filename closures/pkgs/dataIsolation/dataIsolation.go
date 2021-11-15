package dataisolation

import (
	"fmt"
	"net/http"
)

type Database struct {
	Name string
	Url  string
}

func NewDatabase(name, url string) Database {
	return Database{
		Name: name,
		Url:  url,
	}
}

func DatabaseRouteGen(db Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Database: %s\nUrl: %s\n", db.Name, db.Url)
	}
}
