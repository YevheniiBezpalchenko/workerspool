package repository

import (
	"database/sql"
	"log"
	"workerspool/models"
)

type ProductDB struct {
	conn *sql.DB
}

func (r *ProductDB) Conn(conn *sql.DB) {
	r.conn = conn
}

func (db *ProductDB) Create(menu models.Menu) (id int64, err error) {
	stmt, err := db.conn.Prepare("INSERT TO Product(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(menu.Name)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastId, err
}
