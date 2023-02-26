package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Load_Single tests loading a single resource.
func Test_Load_Single(t *testing.T) {
	t.Parallel()
	resources, err := LoadFile("testdata/single.yaml")
	assert.NoError(t, err)
	assert.Len(t, resources, 1)
	assert.Equal(t, "name", resources[0].Name)
	assert.Equal(t, "description", resources[0].Description)
	assert.Equal(t, "owner", resources[0].Owner)
	assert.Equal(t, "resource", resources[0].Kind)
	assert.Equal(t, "test", resources[0].Data.(string))
}

// Test_Load_Multiple tests loading multiple resources.
func Test_Load_Multiple(t *testing.T) {
	t.Parallel()
	resources, err := LoadFile("testdata/multiple.yaml")
	assert.NoError(t, err)
	assert.Len(t, resources, 3)
	assert.Equal(t, "name", resources[0].Name)
	assert.Equal(t, "description", resources[0].Description)
	assert.Equal(t, "owner", resources[0].Owner)
	assert.Equal(t, "resource", resources[0].Kind)
	assert.Equal(t, "test", resources[0].Data.(string))
	assert.Equal(t, "name2", resources[1].Name)
	assert.Equal(t, "description2", resources[1].Description)
	assert.Equal(t, "owner2", resources[1].Owner)
	assert.Equal(t, "resource2", resources[1].Kind)
	assert.Equal(t, "test2", resources[1].Data.(string))
	assert.Equal(t, "name3", resources[2].Name)
	assert.Equal(t, "description3", resources[2].Description)
	assert.Equal(t, "owner3", resources[2].Owner)
	assert.Equal(t, "resource3", resources[2].Kind)
	assert.Equal(t, map[string]interface{}{
		"key": "value",
		"array": []interface{}{
			"value1",
		},
	}, resources[2].Data)
}

// Test_Load_Corrupted tests loading a corrupted resource.
func Test_Load_Corrupted(t *testing.T) {
	t.Parallel()
	_, err := LoadFile("testdata/corrupted.yaml")
	assert.Error(t, err)
}

// Test_Load_Missing tests loading a resource with missing fields.
func Test_Load_Missing(t *testing.T) {
	t.Parallel()
	r, err := LoadFile("testdata/missing.yaml")
	assert.NoError(t, err)
	assert.Len(t, r, 1)
	assert.Equal(t, "", r[0].Name)
	assert.Equal(t, "", r[0].Description)
	assert.Equal(t, "", r[0].Owner)
}

// Test_Load_Empty tests loading an empty resource.
func Test_Load_Empty(t *testing.T) {
	t.Parallel()
	r, err := LoadFile("testdata/empty.yaml")
	assert.NoError(t, err)
	assert.Len(t, r, 0)
}

// Test_Load_Nonexistent tests loading a nonexistent resource.
func Test_Load_Nonexistent(t *testing.T) {
	t.Parallel()
	_, err := LoadFile("testdata/nonexistent.yaml")
	assert.Error(t, err)
}
