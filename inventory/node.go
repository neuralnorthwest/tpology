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
	"github.com/neuralnorthwest/tpology/resource"
)

// Edge is a relationship between a source and a target. The source is the
// resource that has an explicit relationship with the target. The target is
// the resource that is being related to.
type Edge struct {
	// Source is the source node.
	Source *Node
	// Target is the target node.
	Target *Node
}

// Node is a node in the resource graph. It contains a resource, an incoming
// edge map, and an outgoing edge map.
type Node struct {
	// Resource is the resource.
	*resource.Resource
	// Incoming is the incoming edge map.
	Incoming map[string]*Edge
	// Outgoing is the outgoing edge map.
	Outgoing map[string]*Edge
}

// NewNode returns a new node.
func NewNode(r *resource.Resource) *Node {
	return &Node{
		Resource: r,
		Incoming: make(map[string]*Edge),
		Outgoing: make(map[string]*Edge),
	}
}

// AddIncomingEdge adds an incoming edge to the node.
func (n *Node) AddIncomingEdge(name string, e *Edge) {
	n.Incoming[name] = e
}

// AddOutgoingEdge adds an outgoing edge to the node.
func (n *Node) AddOutgoingEdge(name string, e *Edge) {
	n.Outgoing[name] = e
}
