package repository

import (
	"workerspool/models"
)

type RestRepositoryInterface interface {
	Create(user *models.Rest) int
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
	Create(rest *models.Rest, product *models.Products) int
}
