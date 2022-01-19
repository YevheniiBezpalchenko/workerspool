package repository

import (
	"database/sql"
	"log"
	models "workerspool/models"
)

type prodIngredDB struct {
	conn *sql.DB
}

func (r *prodIngredDB) Conn(conn *sql.DB) {
	r.conn = conn
}

func (r *prodIngredDB) Create(ingredient *models.Ingredients, product *models.Products) (id int64, err error) {
	stmt, err := r.conn.Prepare("INSERT TO Product_Ingredients(product_id,ingredient_id) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(product.Id, ingredient.Id)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastId, err
}
