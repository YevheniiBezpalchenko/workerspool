package repository

import (
	"database/sql"
	"log"
	models "workerspool/models"
)

type restProdDB struct {
	conn *sql.DB
}

func (r *restProdDB) Conn(conn *sql.DB) {
	r.conn = conn
}

func (r *restProdDB) Create(rest *models.Rest, product *models.Products, price float32) (id int64, err error) {
	stmt, err := r.conn.Prepare("INSERT TO Rest_Products(product_id, rest_id, price) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(rest.Id, product.Id, price)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastId, err
}
