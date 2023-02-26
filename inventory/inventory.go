package inventory

import "github.com/neuralnorthwest/tpology/resource"

// Inventory is the inventory.
type Inventory struct {
	// Resources are the resources organized by kind and name.
	Resources map[string]map[string]*resource.Resource
}

// New returns a new inventory.
func New() *Inventory {
	return &Inventory{
		Resources: make(map[string]map[string]*resource.Resource),
	}
}

// AddResource adds a resource to the inventory.
func (inv *Inventory) AddResource(r *resource.Resource) {
	if inv.Resources[r.Kind] == nil {
		inv.Resources[r.Kind] = make(map[string]*resource.Resource)
	}
	inv.Resources[r.Kind][r.Name] = r
}
