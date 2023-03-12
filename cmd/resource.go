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
	List  ResourceListConfig
	Graph ResourceGraphConfig
}

// SetupFlags sets up the flags for the resource command.
func (c *ResourceConfig) SetupFlags(cmd *cobra.Command) {
}

// resourceCommand returns the resource command.
func resourceCommand(config *Config) *cobra.Command {
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
	cmd.AddCommand(resourceListCommand(config))
	cmd.AddCommand(resourceGraphCommand(config))
	return cmd
}
