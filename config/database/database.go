package database

type Database struct {
	address  string
	port     uint
	username string
	password string
	dbName   string
	path     string
	dsn      string
}

func New() *Database {
	return &Database{}
}

func Default() *Database {
	return &Database{
		address: "localhost",
		port:    5432,
		path:    "./",
		dsn:     "gothic_db",
	}
}

func (db *Database) Username() string {
	return db.username
}

func (db *Database) Password() string {
	return db.password
}

func (db *Database) Address() string {
	return db.address
}

func (db *Database) Port() uint {
	return db.port
}

func (db *Database) DSN() string {
	return db.dsn
}

func (db *Database) IsDSNEmpty() bool {
	return db.dsn == ""
}

func (db *Database) IsValid() bool {
	return (db.address == "") || db.username == "" || db.password == "" || db.dbName == ""
}
