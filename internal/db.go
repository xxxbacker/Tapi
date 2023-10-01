package internal

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func init() {
	connStr := GenConStr()
	// Подключиться к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Проверить соединение с базой данных
	err = db.Ping()
	if err != nil {
		log.Fatal(err, " no ping")
	} else {
		log.Println("соединенние установлено")
	}
}
