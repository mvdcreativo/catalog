package bootstrap

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/config"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/category"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/db/mongo_db"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/product"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	ProductHandler  *product.ProductHandler
	CategoryHandler *category.CategoryHandler
	MongoClient     *mongo.Client
	Config          *config.Config
}

func InitializeApp() *App {
	cfg := config.LoadConfig()

	mongoClient, err := mongo_db.ConnectDB()
	if err != nil {
		log.Fatalf("❌ Error conectando a MongoDB: %v", err)
	}

	// Repositorios
	productRepo := product.NewProductRepository(mongoClient, cfg.DbName, "products")
	categoryRepo := category.NewCategoryRepository(mongoClient, cfg.DbName, "categories")

	// Servicios
	productService := product.NewProductService(productRepo)
	categoryService := category.NewCategoryService(categoryRepo)

	// Handlers
	productHandler := product.NewProductHandler(productService)
	categoryHandler := category.NewCategoryHandler(categoryService)

	return &App{
		ProductHandler:  productHandler,
		CategoryHandler: categoryHandler,
		MongoClient:     mongoClient,
		Config:          cfg,
	}
}

func (a *App) SetupRouter() *gin.Engine {
	r := gin.Default()

	routes.SetupRoutes(r,
		a.ProductHandler,
		a.CategoryHandler,
	)
	return r
}
