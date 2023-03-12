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
	"testing"

	"github.com/neuralnorthwest/tpology/inventory"
	inventory_test "github.com/neuralnorthwest/tpology/inventory/test"
	"github.com/stretchr/testify/assert"
)

// Test_ResourceGraph tests the resource graph command.
func Test_ResourceGraph(t *testing.T) {
	resources := inventory_test.MakeTestResources()
	inv := inventory.New(resources...)
	outfile, err := os.CreateTemp("", "tpology-resource-graph-test")
	assert.NoError(t, err)
	outfile.Close()
	defer os.Remove(outfile.Name())
	config := &Config{
		Resource: ResourceConfig{
			Graph: ResourceGraphConfig{
				Format: "dot",
				Output: outfile.Name(),
			},
		},
	}
	err = resourceGraph(inv, config)
	assert.NoError(t, err)
	assert.FileExists(t, outfile.Name())
	data, err := os.ReadFile(outfile.Name())
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}
