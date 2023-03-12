// Copyright 2023 Scott M. Long
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inventory

import "github.com/neuralnorthwest/tpology/resource"

// Inventory is the inventory.
type Inventory struct {
	// Resources are the resources organized by kind and name.
	Resources map[string]map[string]*resource.Resource
}

// New returns a new inventory.
func New(resources ...*resource.Resource) *Inventory {
	inv := &Inventory{
		Resources: make(map[string]map[string]*resource.Resource),
	}
	for _, r := range resources {
		inv.AddResource(r)
	}
	return inv
}

// AddResource adds a resource to the inventory.
func (inv *Inventory) AddResource(r *resource.Resource) {
	if inv.Resources[r.Kind] == nil {
		inv.Resources[r.Kind] = make(map[string]*resource.Resource)
	}
	inv.Resources[r.Kind][r.Name] = r
}
