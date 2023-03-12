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
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// GlobalConfig is the global configuration.
type GlobalConfig struct {
	// Quiet suppresses all output.
	Quiet bool
	// GitCacheDir is the path to the Git cache directory.
	GitCacheDir string
	// InventoryLocal is the path to the local inventory repository
	InventoryLocal string
	// Inventory is the URL to the inventory repository
	Inventory string
	// InventoryRef is the Git reference to the inventory repository
	InventoryRef string
}

// Config holds all configuration.
type Config struct {
	// Global is the global configuration.
	Global GlobalConfig
	// Resource is the resource configuration.
	Resource ResourceConfig
}

// SetupFlags sets up the flags for the itool command.
func (c *GlobalConfig) SetupFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&c.Quiet, "quiet", "q", false, "suppress all output")
	cmd.PersistentFlags().StringVar(&c.GitCacheDir, "git-cache-dir", gitCacheDir(), "path to the Git cache directory")
	cmd.PersistentFlags().StringVarP(&c.InventoryLocal, "inventory-local", "l", "", "path to the local inventory repository")
	cmd.PersistentFlags().StringVarP(&c.Inventory, "inventory", "i", "https://github.com/ZeroEyesTech/ZE-Inventory.git", "URL to the inventory repository")
	cmd.PersistentFlags().StringVarP(&c.InventoryRef, "inventory-ref", "r", "main", "Git reference to the inventory repository")
}

// gitCacheDir returns the path to the user's git cache directory.
func gitCacheDir() string {
	cache, err := os.UserCacheDir()
	if err != nil {
		cache = os.TempDir()
	}
	return filepath.Join(cache, "itool", "git")
}
