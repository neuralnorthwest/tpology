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

package inventory

import (
	"fmt"

	"github.com/neuralnorthwest/tpology/resource"
)

// Graph holds the resource graph
type Graph struct {
	// Nodes are the nodes in the graph.
	Nodes map[string]map[string]*Node
}

// Node is a node in the graph.
type Node struct {
	// Resource is the resource.
	*resource.Resource
	// GraphData is the data of the resource, with outgoing edges replaced with
	// pointers to their target nodes.
	GraphData interface{}
}

// BuildGraph builds the graph from the inventory.
func (inv *Inventory) BuildGraph() (*Graph, error) {
	g := &Graph{
		Nodes: make(map[string]map[string]*Node),
	}
	for _, kind := range inv.Resources {
		for _, r := range kind {
			g.addNode(r)
		}
	}
	if err := g.buildGraph(); err != nil {
		return nil, err
	}
	return g, nil
}

// addNode adds a node to the graph.
func (g *Graph) addNode(r *resource.Resource) {
	if g.Nodes[r.Kind] == nil {
		g.Nodes[r.Kind] = make(map[string]*Node)
	}
	node := &Node{
		Resource: r,
	}
	g.Nodes[r.Kind][r.Name] = node
}

// BuildGraph builds the graph.
func (g *Graph) buildGraph() error {
	for _, kind := range g.Nodes {
		for _, node := range kind {
			switch data := node.Data.(type) {
			case map[string]interface{}:
				// Copy the data and replace any outgoing edges with pointers to
				// their target nodes.
				graphData := make(map[string]interface{})
				if err := g.buildGraphData(node, data, graphData); err != nil {
					return err
				}
				node.GraphData = graphData
			default:
				// Just copy the data.
				node.GraphData = data
			}
		}
	}
	return nil
}

// buildGraphData builds the graph data.
func (g *Graph) buildGraphData(node *Node, data, graphData map[string]interface{}) error {
	for key, value := range data {
		switch v := value.(type) {
		case map[string]interface{}:
			// Copy the data and replace any outgoing edges with pointers to
			// their target nodes.
			graphData[key] = make(map[string]interface{})
			if err := g.buildGraphData(node, v, graphData[key].(map[string]interface{})); err != nil {
				return err
			}
		case string:
			// If the value is an edge, replace it with a pointer to the target
			// node.
			if r, ok := g.findNode(key, v); ok {
				graphData[key] = g.Nodes[r.Kind][r.Name]
			} else {
				return fmt.Errorf("in Resource %s/%s, a reference to %s/%s appeared, but that Resource does not exist", node.Kind, node.Name, key, v)
			}
		default:
			// Just copy the data.
			graphData[key] = v
		}
	}
	return nil
}

// findNode finds a node in the graph.
func (g *Graph) findNode(kind, name string) (*Node, bool) {
	if g.Nodes[kind] == nil {
		return nil, false
	}
	if g.Nodes[kind][name] == nil {
		return nil, false
	}
	return g.Nodes[kind][name], true
}
