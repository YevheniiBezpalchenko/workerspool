package repository

import (
	"database/sql"
	"log"
	"workerspool/models"
	"workerspool/models/parser"
)

type RestDB struct {
	conn *sql.DB
}

func (r *RestDB) Conn(conn *sql.DB) {
	r.conn = conn
}

func (r *RestDB) Create(supp *models.Supplier) (id int64, err error) {
	stmt, err := r.conn.Prepare("INSERT TO Suppliers(name,image,open,close,api_id) VALUES(?,?,?,?,?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(supp.Name, supp.Image, supp.Opening, supp.Closing, supp.Api_id)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastId, err
}
func (r *RestDB) FillSupplier(supplier *parser.Suppliers) {
	newsupp := models.Supplier{
		Name:    supplier.Name,
		Image:   supplier.Image,
		Opening: supplier.WorkingHours.Opening,
		Closing: supplier.WorkingHours.Closing,
		Api_id:  supplier.Id,
	}
	r.Create(&newsupp)
}
