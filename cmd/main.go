package main

import (
	"log"

	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/bootstrap"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/watchers"
)

func main() {

	app := bootstrap.InitializeApp()

	r := app.SetupRouter()

	go watchers.WatchCategoryChanges(app.MongoClient, app.Config.DbName)

	log.Printf("🚀 Servidor corriendo en el puerto %s", app.Config.Port)
	if err := r.Run(":" + app.Config.Port); err != nil {
		log.Fatal(err)
	}
}
