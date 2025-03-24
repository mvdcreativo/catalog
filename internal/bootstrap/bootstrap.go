package bootstrap

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/config"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/handlers"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/repositories"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/routes"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	ProductHandler  *handlers.ProductHandler
	CategoryHandler *handlers.CategoryHandler
	MongoClient     *mongo.Client
	Config          *config.Config
}

func InitializeApp() *App {
	cfg := config.LoadConfig()

	mongoClient, err := repositories.ConnectDB()
	if err != nil {
		log.Fatalf("❌ Error conectando a MongoDB: %v", err)
	}

	// Repositorios
	productRepo := repositories.NewProductRepository(mongoClient, cfg.DbName, "products")
	categoryRepo := repositories.NewCategoryRepository(mongoClient, cfg.DbName, "categories")

	// Servicios
	productService := services.NewProductService(productRepo)
	categoryService := services.NewCategoryService(categoryRepo)

	// Handlers
	productHandler := handlers.NewProductHandler(productService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	return &App{
		ProductHandler:  productHandler,
		CategoryHandler: categoryHandler,
		MongoClient:     mongoClient,
		Config:          cfg,
	}
}

func (a *App) SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	routes.RegisterProductRoutes(api, a.ProductHandler)
	routes.RegisterCategoryRoutes(api, a.CategoryHandler)

	return r
}
