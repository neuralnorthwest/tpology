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
// limitations under the License.package inventory

import (
	"fmt"

	"github.com/neuralnorthwest/tpology/resource"
)

// buildGraph builds the resource graph.
func (inv *Inventory) buildGraph() error {
	for _, resources := range inv.Resources {
		for _, r := range resources {
			inv.addNode(r)
		}
	}
	for _, node := range inv.Nodes {
		for _, n := range node {
			if err := inv.addEdges(n); err != nil {
				return err
			}
		}
	}
	return nil
}

// addNode adds a node to the graph.
func (inv *Inventory) addNode(r *resource.Resource) {
	if inv.Nodes[r.Kind] == nil {
		inv.Nodes[r.Kind] = make(map[string]*Node)
	}
	inv.Nodes[r.Kind][r.Name] = NewNode(r)
}

// addEdges adds edges to the graph.
func (inv *Inventory) addEdges(n *Node) error {
	// find all the outgoing edges from this node. These are the references
	// from this node to other nodes. A reference exists at any point where
	// a key in the resource is the name of a kind in the inventory.
	references, err := inv.references(n)
	if err != nil {
		return err
	}
}

// reference is a reference within a node
type reference struct {
	// Kind is the kind of the reference.
	Kind string
	// Map is the map that contains the reference.
	Map map[string]interface{}
	// Key is the key of the reference.
	Key string
	// Value is the value of the reference.
	Value string
}

// references returns a list of all outgoing references from a node.
func (inv *Inventory) references(n *Node) (map[string][]string, error) {
	references := make(map[string][]string)
	switch data := n.Data.(type) {
	case map[string]interface{}:
		for k, v := range data {
			if inv.Nodes[k] == nil {
				continue
			}
			switch v := v.(type) {
			case string:
				references[k] = append(references[k], v)
			case []string:
				references[k] = append(references[k], v...)
			default:
				return nil, fmt.Errorf("invalid reference type %T", v)
			}
		}
