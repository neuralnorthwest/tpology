package cmd

import (
	"github.com/neuralnorthwest/tpology/git"
	"github.com/neuralnorthwest/tpology/inventory"
)

// loadInventory loads the inventory.
func loadInventory() (*inventory.Inventory, error) {
	invPath := config.Global.InventoryLocal
	if invPath == "" {
		cache := git.NewCache(config.Global.GitCacheDir)
		invRepo := cache.New(config.Global.Inventory, config.Global.InventoryRef)
		invPath = invRepo.Dir
	}
	return inventory.Load(invPath)
}
