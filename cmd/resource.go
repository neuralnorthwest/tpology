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
	"github.com/spf13/cobra"
)

// ResourceConfig is the resource configuration.
type ResourceConfig struct {
	// List is the resource list configuration.
	List ResourceListConfig
}

// SetupFlags sets up the flags for the resource command.
func (c *ResourceConfig) SetupFlags(cmd *cobra.Command) {
}

// ResourceListConfig is the resource list configuration.
type ResourceListConfig struct {
	// Format is the output format.
	Format string
}

// SetupFlags sets up the flags for the resource list command.
func (c *ResourceListConfig) SetupFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&c.Format, "format", "f", "table", "output format")
}

// resourceCommand returns the resource command.
func resourceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "resource",
		Aliases: []string{"res", "resources"},
		Short:   "Manage resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	config.Resource.SetupFlags(cmd)
	cmd.AddCommand(resourceListCommand())
	return cmd
}

// resourceListCommand returns the resource list command.
func resourceListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list [resource]",
		Aliases: []string{"ls"},
		Short:   "List resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return resourceList(args)
		},
		Args:          cobra.MaximumNArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	config.Resource.List.SetupFlags(cmd)
	return cmd
}

// resourceList lists resources.
func resourceList(args []string) error {
	inv, err := loadInventory()
	if err != nil {
		return err
	}
	ents := []interface{}{}
	if len(args) == 0 {
		for kind := range inv.Resources {
			ents = append(ents, kind)
		}
	} else {
		kind := args[0]
		if resources, ok := inv.Resources[kind]; ok {
			for _, r := range resources {
				ents = append(ents, r)
			}
		}
	}
	return printEntities(ents, Format(config.Resource.List.Format))
}
