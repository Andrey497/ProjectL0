package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin0"
	password = "1234"
	dbname   = "projectl0"
)

type DBConnect struct {
	Db  *sql.DB
	Err error
}

func (d *DBConnect) OpenConnect() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	d.Db, d.Err = sql.Open("postgres", psqlconn)
	if d.Err != nil {
		panic(d.Err)
	}
}

func (d *DBConnect) CloseConnect() {
	d.Db.Close()
}
