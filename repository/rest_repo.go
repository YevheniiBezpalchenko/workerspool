package repository

import (
	"database/sql"
	"log"
	models "workerspool/models"
)

type RestDB struct {
	conn *sql.DB
}

func (r *RestDB) Conn(conn *sql.DB) {
	r.conn = conn
}

func (r *RestDB) Create(rest *models.Rest) (id int64, err error) {
	stmt, err := r.conn.Prepare("INSERT TO Rest(name,image,open,close) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(rest.Name, rest.Image, rest.WorkingHours.Opening, rest.WorkingHours.Closing)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastId, err
}
