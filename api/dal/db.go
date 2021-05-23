package dal

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var once sync.Once

func Connect() (*sql.DB, error) {
	var db *sql.DB
	var err error
	//once.Do(func() {
	// dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", DB_USER, DB_PASSWORD, DB_NAME, SSL_MODE)
	// db, _ = sql.Open("postgres", dbinfo)
	db, _ = sql.Open("sqlite3", "app.db")
	err = db.Ping()
	//})
	return db, err
}

func LogAndQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	fmt.Println(query)
	return db.Query(query, args...)
}

func MustExec(db *sql.DB, query string, args ...interface{}) {
	_, err := db.Exec(query, args...)
	if err != nil {
		panic(err)
	}
}

func LogAndExec(db *sql.DB, query string, args ...interface{}) (int64, error) {
	fmt.Println(query)

	res, err := db.Exec(query, args...)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()

	return id, err
}
