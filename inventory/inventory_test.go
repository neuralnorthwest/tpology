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

	"github.com/neuralnorthwest/tpology/resource"
	"github.com/stretchr/testify/assert"
)

// Test_New tests the inventory new function.
func Test_New(t *testing.T) {
	t.Parallel()
	inv := New()
	assert.NotNil(t, inv)
	assert.Len(t, inv.Resources, 0)
}

// Test_Inventory_AddResource tests the inventory add resource function.
func Test_Inventory_AddResource(t *testing.T) {
	t.Parallel()
	inv := New()
	inv.AddResource(&resource.Resource{
		Kind: "resource",
		Name: "foo",
	})
	assert.Len(t, inv.Resources, 1)
	assert.Len(t, inv.Resources["resource"], 1)
	assert.Equal(t, "foo", inv.Resources["resource"]["foo"].Name)
}
