package product

import (
	"fmt"
	"time"

	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/domains/category"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/interfaces/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product representa un producto en la colección "products".
type Product struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Store            StoreRef           `bson:"store" json:"store"`                                                     // Referencia a la tienda propietaria.
	Name             string             `bson:"name" json:"name" validate:"required,min=3,max=50"`                      // Nombre del producto.
	ShortDescription string             `bson:"short_description" json:"short_description" validate:"required,max=200"` // Descripción breve.
	LongDescription  string             `bson:"long_description" json:"long_description" validate:"max=5000"`           // Descripción detallada.
	SKU              string             `bson:"sku" json:"sku" `                                                        // Código SKU.
	Price            float64            `bson:"price" json:"price" validate:"required"`                                 // Precio base.
	Currency         string             `bson:"currency" json:"currency" validate:"required,min=3,max=3" `              // Moneda (ej. USD, UYU).
	Stock            int                `bson:"stock" json:"stock validate:"number"`                                    // Cantidad en inventario.
	Active           bool               `bson:"active" json:"active"`                                                   // Indica si el producto está activo.
	CreatedAt        time.Time          `bson:"created_at" json:"created_at"`                                           // Fecha de creación.
	UpdatedAt        time.Time          `bson:"updated_at" json:"updated_at"`                                           // Fecha de última actualización.
	Status           Status             `bson:"status" json:"status" validate:"required,status"`                        // Estado del producto.

	// Relación con categorías:
	// Opción B: Snapshot embebido de la información mínima de cada categoría.
	Categories []category.CategoryRefDTO `bson:"categories,omitempty" json:"categories,omitempty"`

	ParentProductID *primitive.ObjectID  `bson:"parent_product_id,omitempty" json:"parent_product_id,omitempty"` // Referencia al producto padre, si es variante.
	Images          []storage.FileObject `bson:"images" json:"images"`                                           // Imágenes asociadas.
	Variants        []ProductVariant     `bson:"variants" json:"variants"`                                       // Variantes u opciones.
}

type Status string

const (
	StatusPending    Status = "PENDING"
	StatusActive     Status = "PUBLISHED"
	StatusOutOfStock Status = "OUT_OF_STOCK"
	StatusArchived   Status = "ARCHIVED"
)

// ValidStatuses devuelve un slice con todos los valores válidos
func ValidStatusesProduct() []string {
	return []string{
		string(StatusPending),
		string(StatusActive),
		string(StatusOutOfStock),
		string(StatusArchived),
	}
}

type StoreRef struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"` // Identificador único de la tienda.
	Name string             `bson:"name" json:"name"`        // Nombre de la tienda.
}

// ProductVariant representa una variante u opción del producto.
type ProductVariant struct {
	Options         Options `bson:"options" json:"options"`                   // Opciones que definen la variante.
	PriceAdjustment float64 `bson:"price_adjustment" json:"price_adjustment"` // Ajuste de precio para esta variante.
	VariantStock    int     `bson:"variant_stock" json:"variant_stock"`       // Stock específico para la variante.
}

// Options define las opciones de la variante.
// Puedes ajustar estos campos o usar un mapa para opciones dinámicas.
type Options struct {
	Color string `bson:"color" json:"color"` // Ejemplo: "Red"
	Size  string `bson:"size" json:"size"`   // Ejemplo: "M"
}

var ProductFilterWhitelist = map[string]bool{
	"name":                   true,
	"price":                  true,
	"currency":               true,
	"stock":                  true,
	"active":                 true,
	"status":                 true,
	"store.id":               true,
	"store.name":             true,
	"categories.name":        true,
	"categories.id":          true,
	"parent_product_id":      true,
	"variants.options.color": true,
	"variants.options.size":  true,
}

func (Product) GetFilterWhitelist() (map[string]bool, error) {
	if ProductFilterWhitelist == nil {
		return nil, fmt.Errorf("filter whitelist is not initialized")
	}
	return ProductFilterWhitelist, nil
}

func (p *Product) SetID(id primitive.ObjectID) {
	p.ID = id
}

func (p *Product) SetCreationDate(t time.Time) {
	p.CreatedAt = t
}

func (p *Product) SetUpdateDate(t time.Time) {
	p.UpdatedAt = t
}
