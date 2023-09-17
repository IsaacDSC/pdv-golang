package settings

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB
var once sync.Once

func DbConn() *sql.DB {
	dsn := "host=localhost user=root password=root dbname=solutionstech port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	once.Do(func() {
		var err error
		if db, err = sql.Open("postgres", dsn); err != nil {
			log.Panic(err)
		}
		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(20)
	})
	return db
}

func DbClose() {
	db.Close()
}
