package database

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	valid := Database{
		address:  "localhost",
		port:     5432,
		username: "root",
		password: "toor",
		dbName:   "gothic_db",
	}
	invalid := Database{
		address:  "localhost",
		port:     5432,
		username: "",
		password: "toor",
		dbName:   "gothic_db",
	}

	if !valid.IsValid() {
		t.Errorf("")
	}

	if invalid.IsValid() {
		t.Errorf("")
	}
}
