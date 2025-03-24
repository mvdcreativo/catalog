package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/utils/mql_request_filter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CRUDRepository define las operaciones CRUD genéricas para cualquier tipo T.
type CRUDRepository[T any] interface {
	Insert(ctx context.Context, entity *T) error
	FindByID(ctx context.Context, id string) (*T, error)
	Update(ctx context.Context, id string, entity *T) error
	Delete(ctx context.Context, filter interface{}) error
	// FindAll(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]T, error)
	FindAll(ctx context.Context, filter map[string]interface{}, page, limit int) ([]T, int64, error)
}

// crudRepository es una implementación genérica de CRUDRepository.
type crudRepository[T any] struct {
	collection *mongo.Collection
}

// NewCRUDRepository crea un nuevo repositorio CRUD genérico para la colección dada.
func NewCRUDRepository[T any](client *mongo.Client, dbName, collectionName string) CRUDRepository[T] {
	col := client.Database(dbName).Collection(collectionName)
	return &crudRepository[T]{
		collection: col,
	}
}

func (r *crudRepository[T]) Insert(ctx context.Context, entity *T) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, entity)
	return err
}

func (r *crudRepository[T]) FindByID(ctx context.Context, id string) (*T, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("ID inválido: %w", err)
	}

	var entity T
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&entity)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *crudRepository[T]) Update(ctx context.Context, id string, entity *T) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("ID inválido: %w", err)
	}

	update := bson.M{"$set": entity}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)

	return err
}

func (r *crudRepository[T]) Delete(ctx context.Context, filter interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}

// func (r *crudRepository[T]) FindAll(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]T, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
// 	defer cancel()

// 	// opts := options.Find().SetSkip(skip).SetLimit(int64(limit))
// 	// cursor, err := r.collection.Find(ctx, bson.M{}, opts)

// 	cursor, err := r.collection.Find(ctx, filter, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)
// 	var results []T
// 	for cursor.Next(ctx) {
// 		var entity T
// 		if err := cursor.Decode(&entity); err != nil {
// 			return nil, err
// 		}
// 		results = append(results, entity)
// 	}
// 	return results, nil
// }

func (r *crudRepository[T]) FindAll(ctx context.Context, filter map[string]interface{}, page, limit int) ([]T, int64, error) {
	mongoFilter := bson.M{}
	for k, v := range filter {
		mongoFilter[k] = v // Soporta dot notation
	}

	opts := options.Find().
		SetSkip(int64(mql_request_filter.GetOffset(page, limit))).
		SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, mongoFilter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var results []T
	if err := cursor.All(ctx, &results); err != nil {
		return nil, 0, err
	}

	// Obtener el total de documentos en la colección
	total, err := r.collection.CountDocuments(ctx, mongoFilter)
	if err != nil {
		return nil, 0, err
	}

	return results, total, nil
}
