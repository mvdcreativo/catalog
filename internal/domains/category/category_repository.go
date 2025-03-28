package category

import (
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/generics/repository/mongo_repository"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/interfaces/i_crud"
	"go.mongodb.org/mongo-driver/mongo"
)

// CategoryRepository define las operaciones de acceso a datos para Category.
type CategoryRepository interface {
	i_crud.CRUDRepository[Category] // Hereda los métodos genéricosgenéricos.

	// Insert(category *category.Category) error ejemplo de un custom repo

}

type categoryRepository struct {
	i_crud.CRUDRepository[Category]
}

// NewCategoryRepository crea un nuevo repositorio de Category usando la colección "categorys" de la base de datos.
func NewCategoryRepository(client *mongo.Client, dbName, collectionName string) CategoryRepository {
	genRepo := mongo_repository.NewCRUDRepository[Category](client, dbName, collectionName)
	return &categoryRepository{CRUDRepository: genRepo}
}
