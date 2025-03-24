package order

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name             string             `json:"name" bson:"name"`
  CreatedAt        time.Time          `bson:"created_at" json:"created_at"`               // Fecha de creación.
	UpdatedAt        time.Time          `bson:"updated_at" json:"updated_at"`     
}


var OrderFilterWhitelist = map[string]bool{
	"name":                   true,
  "created_at"              true,
  "updated_at"              true,
  
}

func (Order) GetFilterWhitelist() (map[string]bool, error) {
	if OrderFilterWhitelist == nil {
		return nil, fmt.Errorf("filter whitelist is not initialized")
	}
	return OrderFilterWhitelist, nil
}

func (p *Order) SetID(id primitive.ObjectID) {
	p.ID = id
}

func (p *Order) SetCreationDate(t time.Time) {
	p.CreatedAt = t
}

func (p *Order) SetUpdateDate(t time.Time) {
	p.UpdatedAt = t
}
