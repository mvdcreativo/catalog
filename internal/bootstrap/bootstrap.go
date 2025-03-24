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

// App contiene todas las dependencias que se inyectan en la aplicación
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

// InitializeApp carga la configuración, conecta a la base de datos y registra todos los módulos
func InitializeApp() *App {
	cfg := config.LoadConfig()

	client, err := mongo_db.ConnectDB()
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

	return app
}

// SetupRouter crea y devuelve un router Gin con las rutas configuradas
func (app *App) SetupRouter() *gin.Engine {
	r := gin.Default()
	routes.SetupRoutes(r, app.ProductHandler, app.CategoryHandler)
	return r
}

// --- REGISTRO DE MÓDULOS ---

func registerProduct(app *App) {
	repo := product.NewProductRepository(app.MongoClient, app.Config.DbName, "products")
	service := product.NewProductService(repo)
	app.ProductHandler = product.NewProductHandler(service)
}

func registerCategory(app *App) {
	repo := category.NewCategoryRepository(app.MongoClient, app.Config.DbName, "categories")
	service := category.NewCategoryService(repo)
	app.CategoryHandler = category.NewCategoryHandler(service)
}
