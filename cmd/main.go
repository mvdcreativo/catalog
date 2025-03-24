package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/config"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/handlers"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/repositories"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/routes"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/services"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/watchers"
)

func init() {
	// Conectar a MongoDB
	repositories.ConnectDB()
}

func main() {
	// Cargar configuración
	c := config.LoadConfig()

	// Products
	productRepo := repositories.NewProductRepository(repositories.MongoClient)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Categories
	categoryRepo := repositories.NewCategoryRepository(repositories.MongoClient)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Crear instancia de Gin
	r := gin.Default()
	// Definir rutas
	routes.SetupRoutes(
		r,
		productHandler,
		categoryHandler,
	)

	go watchers.WatchCategoryChanges(repositories.MongoClient, c.DbName)

	// Iniciar servidor
	port := c.Port
	log.Printf("Servidor corriendo en el puerto %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
