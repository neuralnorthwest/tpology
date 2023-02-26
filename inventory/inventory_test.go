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
