package order

import (
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/db/mongo_db/mongo_repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	mongo_repository.CRUDRepository[Order] 
}

type orderRepository struct {
	mongo_repository.CRUDRepository[Order]
}

func NewOrderRepository(client *mongo.Client, dbName, collectionName string) OrderRepository {
	genRepo := mongo_repository.NewCRUDRepository[Order](client, dbName, collectionName)
	return &orderRepository{CRUDRepository: genRepo}
}
