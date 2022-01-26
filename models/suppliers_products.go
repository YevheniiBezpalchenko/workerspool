package models

import "workerspool/models/parser"

type SuppliersProducts struct {
	Id           int
	Product      Products
	Supplier     parser.Suppliers
	Price        float32
	SuppApiId    int
	ProductApiId int
}
