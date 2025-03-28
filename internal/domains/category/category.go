package category

import (
	"fmt"
	"time"

	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/interfaces/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Category representa una categoría en la colección "categories".
// Atributos:
// - Name: Nombre de la categoría.
// - CategoryParentID: Referencia a la categoría padre (si existe).
// - Images: Lista de URLs de imágenes asociadas a la categoría.
// - Icon: URL o nombre del icono de la categoría.
type Category struct {
	ID               primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name             string               `bson:"name" json:"name" validate:"required,min=3,max=50"`
	CategoryParentID *primitive.ObjectID  `bson:"category_parent_id,omitempty" json:"category_parent_id,omitempty" `
	Images           []storage.FileObject `bson:"images,omitempty" json:"images,omitempty"`
	Icon             string               `bson:"icon,omitempty" json:"icon,omitempty" validate:"url"`
	CreatedAt        time.Time            `bson:"created_at" json:"created_at"` // Fecha de creación.
	UpdatedAt        time.Time            `bson:"updated_at" json:"updated_at"` // Fecha de última actualización.
}

// CategoryView es una versión reducida de Category sin el campo Images.
type CategoryRefDTO struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name"`
	Icon string             `bson:"icon,omitempty" json:"icon,omitempty"`
}

// ToBsonM convierte una instancia de Category a un mapa BSON,
// incluyendo únicamente los campos definidos en el struct.
func (c *CategoryRefDTO) ToBsonM() (bson.M, error) {
	data, err := bson.Marshal(c)
	if err != nil {
		return nil, err
	}
	var result bson.M
	if err := bson.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

var CategoryFilterWhitelist = map[string]bool{
	"id":                 false,
	"name":               true,
	"category_parent_id": true,
	"icon":               true,
	"images":             false,
}

func (Category) GetFilterWhitelist() (map[string]bool, error) {
	if CategoryFilterWhitelist == nil {
		return nil, fmt.Errorf("filter whitelist is not initialized")
	}
	return CategoryFilterWhitelist, nil
}

func (i *Category) SetID(id primitive.ObjectID) {
	i.ID = id
}

func (i *Category) SetCreationDate(t time.Time) {
	i.CreatedAt = t
}

func (i *Category) SetUpdateDate(t time.Time) {
	i.UpdatedAt = t
}
