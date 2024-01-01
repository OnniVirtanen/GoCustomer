// Package aggregates holds aggregates that combines many entities into a full object
package aggregate

import (
	"example.com/backend/core/model/entity"
)

// Customer is a aggregate that combines all entities needed to represent a customer
type Customer struct {
	// person is the root entity of a customer
	// which means the person.ID is the main identifier for this aggregate
	Person *entity.Person `json:"person"`
}
