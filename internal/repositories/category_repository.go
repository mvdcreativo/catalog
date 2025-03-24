package repositories

import (
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// CategoryRepository define las operaciones de acceso a datos para Category.
type CategoryRepository interface {
	CRUDRepository[models.Category] // Hereda los métodos genéricosgenéricos.

	// Insert(category *models.Category) error ejemplo de un custom repo

}

type categoryRepository struct {
	CRUDRepository[models.Category]
}

// NewCategoryRepository crea un nuevo repositorio de Category usando la colección "categorys" de la base de datos.
func NewCategoryRepository(client *mongo.Client, dbName, collectionName string) CategoryRepository {
	genRepo := NewCRUDRepository[models.Category](client, dbName, collectionName)
	return &categoryRepository{CRUDRepository: genRepo}
}
