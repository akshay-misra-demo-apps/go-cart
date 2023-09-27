package interfaces

import (
	"github.com/akshay-misra-demo-apps/go-cart/user-management/models"
)

type IProduct interface {
	Create(*models.Product) (*models.Product, error)
	Get(string) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	Search(string) ([]*models.Product, error)
	Patch(*models.Product) (*models.Product, error)
	Delete(string) error
}
