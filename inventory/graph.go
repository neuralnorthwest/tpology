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
			graphData, err := g.buildGraphData(node.Data)
			if err != nil {
				return fmt.Errorf("%s: error building graph data for %s/%s: %w", node.Filename(), node.Kind, node.Name, err)
			}
			node.GraphData = graphData
		}
	}
	return nil
}

// buildGraphData builds the graph data by processing the incoming data
// structure and replacing any references to other resources with pointers to
// the target nodes.
func (g *Graph) buildGraphData(data interface{}) (interface{}, error) {
	switch v := data.(type) {
	case map[string]interface{}:
		return g.buildGraphDataMap(v)
	case []interface{}:
		return g.buildGraphDataList(v)
	default:
		return v, nil
	}
}

// buildGraphDataMap builds the graph data map.
func (g *Graph) buildGraphDataMap(data map[string]interface{}) (map[string]interface{}, error) {
	graphData := make(map[string]interface{})
	for k, v := range data {
		switch v := v.(type) {
		case string:
			if g.Nodes[k] != nil {
				if node, ok := g.Nodes[k][v]; ok {
					graphData[k] = node
					continue
				}
				return nil, fmt.Errorf("resource not found: %s/%s", k, v)
			}
			graphData[k] = v
		default:
			newData, err := g.buildGraphData(v)
			if err != nil {
				return nil, err
			}
			graphData[k] = newData
		}
	}
	return graphData, nil
}

// buildGraphDataList builds the graph data list.
func (g *Graph) buildGraphDataList(data []interface{}) ([]interface{}, error) {
	graphData := make([]interface{}, len(data))
	for i, v := range data {
		newData, err := g.buildGraphData(v)
		if err != nil {
			return nil, err
		}
		graphData[i] = newData
	}
	return graphData, nil
}
