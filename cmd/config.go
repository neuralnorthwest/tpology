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

// config is the global configuration.
var config = &Config{}

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
