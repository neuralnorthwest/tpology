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

package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/neuralnorthwest/tpology/inventory"
	"github.com/spf13/cobra"
)

// ResourceGraphConfig is the resource graph configuration.
type ResourceGraphConfig struct {
	// Format is the output format.
	Format string
	// Output is the output file.
	Output string
}

// SetupFlags sets up the flags for the resource graph command.
func (c *ResourceGraphConfig) SetupFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&c.Format, "format", "f", "dot", "output format")
	cmd.Flags().StringVarP(&c.Output, "output", "o", "", "output file")
}

// resourceGraphCommand returns the resource graph command.
func resourceGraphCommand(config *Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "graph",
		Short: "Graph resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			inv, err := loadInventory(config)
			if err != nil {
				return err
			}
			return resourceGraph(inv, config)
		},
		Args:          cobra.NoArgs,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	config.Resource.Graph.SetupFlags(cmd)
	return cmd
}

// resourceGraph graphs resources.
func resourceGraph(inv *inventory.Inventory, config *Config) error {
	g, err := inv.BuildGraph()
	if err != nil {
		return err
	}
	var w io.Writer = os.Stdout
	if config.Resource.Graph.Output != "" {
		f, err := os.Create(config.Resource.Graph.Output)
		if err != nil {
			return err
		}
		defer f.Close()
		w = f
	}
	switch config.Resource.Graph.Format {
	case "dot":
		return dumpDot(w, g)
	default:
		return fmt.Errorf("unknown format: %s", config.Resource.Graph.Format)
	}
}

// dumpDot dumps the graph in dot format.
func dumpDot(w io.Writer, g *inventory.Graph) error {
	fmt.Fprintln(w, "digraph {")
	for _, nodes := range g.Nodes {
		for _, node := range nodes {
			dumpEdges(w, node, node.GraphData)
		}
	}
	fmt.Fprintln(w, "}")
	return nil
}

// dumpEdges dumps the edges for a node.
func dumpEdges(w io.Writer, node *inventory.Node, data interface{}) {
	switch data := data.(type) {
	case []interface{}:
		for _, elem := range data {
			dumpEdges(w, node, elem)
		}
	case map[string]interface{}:
		for _, elem := range data {
			dumpEdges(w, node, elem)
		}
	case *inventory.Node:
		fmt.Fprintf(w, "    \"%s\" -> \"%s\";\n", node.QualifiedName(), data.QualifiedName())
	}
}
