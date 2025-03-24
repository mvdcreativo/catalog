package product

import (
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/db/mongo_db/mongo_repository"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductRepository define las operaciones de acceso a datos para Product.
type ProductRepository interface {
	mongo_repository.CRUDRepository[Product] // Hereda los métodos genéricosgenéricos.

	// Insert(product *product.Product) error ejemplo de un custom repo

}

type productRepository struct {
	mongo_repository.CRUDRepository[Product]
}

// NewProductRepository crea un nuevo repositorio de Product usando la colección "products" de la base de datos.
func NewProductRepository(client *mongo.Client, dbName, collectionName string) ProductRepository {
	genRepo := mongo_repository.NewCRUDRepository[Product](client, dbName, collectionName)
	return &productRepository{CRUDRepository: genRepo}
}
