package repository

import (
	"database/sql"
	"log"
	models "workerspool/models"
)

type IngredientDB struct {
	conn *sql.DB
}

func (r *IngredientDB) Conn(conn *sql.DB) {
	r.conn = conn
}

func (r *IngredientDB) Create(ingredient *models.Ingredients) (id int64, err error) {
	stmt, err := r.conn.Prepare("INSERT TO Ingredients(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(ingredient.Name)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastId, err
}
