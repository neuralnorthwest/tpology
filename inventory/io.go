package inventory

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/neuralnorthwest/tpology/resource"
)

// Load loads the inventory.
func Load(path string) (*Inventory, error) {
	inv := New()
	if err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".yaml" && ext != ".yml" {
			return nil
		}
		return inv.LoadResourceFile(path)
	}); err != nil {
		return nil, err
	}
	return inv, nil
}

// LoadResourceFile loads resources from a manifest.
func (inv *Inventory) LoadResourceFile(path string) error {
	resources, err := resource.LoadFile(path)
	if err != nil {
		return err
	}
	for _, r := range resources {
		inv.AddResource(r)
	}
	return nil
}
