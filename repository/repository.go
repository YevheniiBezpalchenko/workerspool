package repository

import (
	"workerspool/models"
	"workerspool/models/parser"
)

type RestRepositoryInterface interface {
	Create(user *parser.Rest) int
}
type IngredientRepositoryInterface interface {
	Create(user *models.Ingredients) int
}
type ProductsRepositoryInterface interface {
	Create(product *models.Products) int
}
type ProductIngredientRepositoryInterface interface {
	Create(ingredient *models.Ingredients, product *models.Products) int
}
type RestProductRepositoryInterface interface {
	Create(rest *parser.Rest, product *models.Products) int
}
