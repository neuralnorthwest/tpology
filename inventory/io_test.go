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
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Inventory_Load tests the inventory load function.
func Test_Inventory_Load(t *testing.T) {
	t.Parallel()
	inv, err := Load("testdata")
	assert.Nil(t, err)
	assert.NotNil(t, inv)
	assert.Equal(t, 3, len(inv.Resources))
	assert.Equal(t, 1, len(inv.Resources["resource"]))
	assert.Equal(t, 1, len(inv.Resources["resource2"]))
	assert.Equal(t, 1, len(inv.Resources["resource3"]))
}

// Test_Inventory_Load_Nonexistent tests the inventory load function with a
// nonexistent directory.
func Test_Inventory_Load_Nonexistent(t *testing.T) {
	t.Parallel()
	_, err := Load("nonexistent")
	assert.NotNil(t, err)
}

// Test_Inventory_LoadResourceFile tests the inventory load resource file
// function.
func Test_Inventory_LoadResourceFile(t *testing.T) {
	t.Parallel()
	inv := New()
	err := inv.LoadResourceFile("testdata/single.yaml")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(inv.Resources))
	assert.Equal(t, 1, len(inv.Resources["resource"]))
}

// Test_Inventory_LoadResourceFile_Nonexistent tests the inventory load resource
// file function with a nonexistent file.
func Test_Inventory_LoadResourceFile_Nonexistent(t *testing.T) {
	t.Parallel()
	inv := New()
	err := inv.LoadResourceFile("nonexistent")
	assert.NotNil(t, err)
}
