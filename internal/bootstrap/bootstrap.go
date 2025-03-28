package bootstrap

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/config"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/domains/category"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/domains/product"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/infrastructure/db/mongo_db"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/infrastructure/storage/minio"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/middleware"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	ProductHandler  *product.ProductHandler
	CategoryHandler *category.CategoryHandler
	MongoClient     *mongo.Client
	Config          *config.Config
}

// Módulos registrados: cada uno construye su repositorio, servicio y handler
var modules = []func(app *App){
	registerProduct,
	registerCategory,
}

func InitializeApp() *App {
	cfg := config.LoadConfig()

	client, err := mongo_db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("❌ Error conectando a MongoDB: %v", err)
	}

	app := &App{
		Config:      cfg,
		MongoClient: client,
	}

	for _, register := range modules {
		register(app)
	}

	middleware.Init()

	return app
}

func (app *App) SetupRouter() *gin.Engine {
	r := gin.Default()

	routes.SetupRoutes(r,
		app.ProductHandler,
		app.CategoryHandler,
	)

	return r
}

// --- REGISTRO DE MÓDULOS ---

func registerProduct(app *App) {
	cfg := app.Config

	// Crear cliente MinIO y uploader
	minioClient := minio.NewMinioClient(
		cfg.Bucket.Endpoint,
		cfg.Bucket.Key,
		cfg.Bucket.Secret,
		cfg.Bucket.Name,
		cfg.Bucket.BaseURL,
		cfg.Bucket.UseSSL,
	)

	uploader := minio.NewUploader(minioClient)
	deleter := minio.NewDeleter(minioClient)

	repo := product.NewProductRepository(app.MongoClient, cfg.Database.Name, "products")
	service := product.NewProductService(repo, uploader, deleter)
	handler := product.NewProductHandler(cfg, service)

	app.ProductHandler = handler
}

func registerCategory(app *App) {
	cfg := app.Config

	repo := category.NewCategoryRepository(app.MongoClient, cfg.Database.Name, "categories")
	service := category.NewCategoryService(repo)
	app.CategoryHandler = category.NewCategoryHandler(service)
}
