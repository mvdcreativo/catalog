package watchers

import (
	"context"
	"log"

	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/domains/category"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// WatchCategoryChanges escucha los cambios en la colección "categories"
// y actualiza los snapshots embebidos en la colección "products".
func WatchCategoryChanges(client *mongo.Client, dbName string) {
	ctx := context.Background()
	catCollection := client.Database(dbName).Collection("categories")
	prodCollection := client.Database(dbName).Collection("products")

	// Definir un pipeline para filtrar solo las operaciones "update" y "replace"
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "operationType", Value: bson.D{{Key: "$in", Value: bson.A{"update", "replace"}}}},
		}}},
	}

	// Con la opción UpdateLookup obtenemos el documento completo posterior a la actualización.
	opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)
	stream, err := catCollection.Watch(ctx, pipeline, opts)
	if err != nil {
		log.Fatalf("Error iniciando change stream: %v", err)
	}
	defer stream.Close(ctx)

	log.Println("Change stream de 'categories' iniciado...")

	// Procesar eventos del change stream.
	for stream.Next(ctx) {
		var event bson.M
		if err := stream.Decode(&event); err != nil {
			log.Printf("Error decodificando evento: %v", err)
			continue
		}

		opType, ok := event["operationType"].(string)
		if !ok {
			continue
		}
		log.Printf("Evento detectado: %s", opType)

		if opType == "update" || opType == "replace" {
			fullDoc, ok := event["fullDocument"].(bson.M)
			if !ok {
				log.Println("No se encontró fullDocument en el evento")
				continue
			}

			// Convertir el fullDocument a una instancia de CategoryRefDTO.
			var categoryRefDTO category.CategoryRefDTO
			fullDocBytes, err := bson.Marshal(fullDoc)
			if err != nil {
				log.Printf("Error al marshal fullDoc: %v", err)
				continue
			}
			if err := bson.Unmarshal(fullDocBytes, &categoryRefDTO); err != nil {
				log.Printf("Error al unmarshal fullDoc a CategoryRefDTO: %v", err)
				continue
			}

			// Convertir la instancia CategoryRefDTO a un mapa BSON usando ToBsonM().
			catMap, err := categoryRefDTO.ToBsonM()
			if err != nil {
				log.Printf("Error convirtiendo CategoryRefDTO a bson.M: %v", err)
				continue
			}

			// Extraer el ObjectID de la categoría.
			catID, ok := fullDoc["_id"].(primitive.ObjectID)
			if !ok {
				log.Println("Error convirtiendo _id a ObjectID")
				continue
			}

			// Construir la operación de actualización que reemplaza el elemento del array "categories".
			update := bson.M{
				"$set": bson.M{
					"categories.$[elem]": catMap,
				},
			}

			// Configurar arrayFilters para actualizar solo el elemento con _id == catID.
			updateOpts := options.Update().SetArrayFilters(options.ArrayFilters{
				Filters: []interface{}{bson.M{"elem._id": catID}},
			})

			// Ejecutar updateMany en la colección "products".
			res, err := prodCollection.UpdateMany(ctx, bson.M{"categories._id": catID}, update, updateOpts)
			if err != nil {
				log.Printf("Error actualizando productos para la categoría %s: %v", catID.Hex(), err)
				continue
			}
			log.Printf("Actualizados %d productos para la categoría %s", res.ModifiedCount, catID.Hex())
		}
	}

	if err := stream.Err(); err != nil {
		log.Printf("Error en el change stream: %v", err)
	}
}
