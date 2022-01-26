package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"workerspool/db"
	"workerspool/models/parser"
	"workerspool/repository"
)

func main() {
	suppliers := getSuppliers()
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
}
func do(conn *sql.DB, supplier parser.Suppliers) {
	var r repository.RestDB
	r.Conn(conn)
	r.FillSupplier(&supplier)

}
func getSuppliers() []parser.Suppliers {
	res, err := http.Get("http://foodapi.true-tech.php.nixdev.co/suppliers")
	if err != nil {
		log.Fatal(err)
	}

	var data []parser.Suppliers
	suplMap := make(map[string][]parser.Suppliers)

	json.NewDecoder(res.Body).Decode(&suplMap)
	data = suplMap["suppliers"]
	return data
}
func getSupplierMenu(Id int) []parser.Menu {
	url := "http://foodapi.true-tech.php.nixdev.co/suppliers/" + strconv.Itoa(Id) + "/menu"
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var data []parser.Menu
	menuMap := make(map[string][]parser.Menu)

	json.NewDecoder(res.Body).Decode(&menuMap)
	data = menuMap["menu"]
	fmt.Println(data)
	return data
}

func getOneSupplierMenu(Supplier_Id int, Menu_Id int) parser.Menu {
	url := "http://foodapi.true-tech.php.nixdev.co/suppliers/" + strconv.Itoa(Supplier_Id) + "/menu/" + strconv.Itoa(Menu_Id)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var data parser.Menu

	json.NewDecoder(res.Body).Decode(&data)
	fmt.Println(data)
	return data
}
