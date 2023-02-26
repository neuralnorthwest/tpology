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
