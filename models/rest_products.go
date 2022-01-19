package models

import "workerspool/models/parser"

type Rest_Products struct {
	Id      int
	Product Products
	Rest    parser.Rest
	price   float32
}
