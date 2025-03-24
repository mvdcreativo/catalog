package repositories

import (
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductRepository define las operaciones de acceso a datos para Product.
type ProductRepository interface {
	CRUDRepository[models.Product] // Hereda los métodos genéricosgenéricos.

	// Insert(product *models.Product) error ejemplo de un custom repo

}

type productRepository struct {
	CRUDRepository[models.Product]
}

// NewProductRepository crea un nuevo repositorio de Product usando la colección "products" de la base de datos.
func NewProductRepository(client *mongo.Client) ProductRepository {
	genRepo := NewCRUDRepository[models.Product](client, "products")
	return &productRepository{CRUDRepository: genRepo}
}
