package repositories

import (
	"context"
	"log"
	"time"

	"github.com/mvdcreativo/e-commerce-saas/catalog/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

// ConnectDB inicializa la conexión a MongoDB
func ConnectDB() {
	cfg := config.LoadConfig()
	uri := config.GetDBURI()
	dbName := cfg.DbName

	// Configurar cliente con tiempo de espera
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conectar a MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("❌ Error conectando a MongoDB: %v", err)
	}

	// Verificar conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a MongoDB: %v", err)
	}

	// Asignar cliente y base de datos
	MongoClient = client
	MongoDB = client.Database(dbName)

	log.Println("✅ Conectado a MongoDB exitosamente!")
}

// GetCollection obtiene una colección específica de la base de datos
func GetCollection(collectionName string) *mongo.Collection {
	return MongoDB.Collection(collectionName)
}
