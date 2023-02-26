package resource

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

// Test_New tests the New function.
func Test_New(t *testing.T) {
	t.Parallel()
	_, err := New("kind", "name", "description", "owner")
	assert.NoError(t, err)
}

// Test_New_InvalidKind tests the New function with an invalid kind.
func Test_New_InvalidKind(t *testing.T) {
	reserved := []string{
		"name",
		"description",
		"owner",
	}
	for _, kind := range reserved {
		_, err := New(kind, "name", "description", "owner")
		assert.Error(t, err)
	}
}

// Test_MarshalYAML tests the MarshalYAML function.
func Test_MarshalYAML(t *testing.T) {
	t.Parallel()
	r, err := New("kind", "name", "description", "owner")
	assert.NoError(t, err)
	r.Data = map[string]interface{}{
		"key": "value",
	}
	data, err := r.MarshalYAML()
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"name":        "name",
		"description": "description",
		"owner":       "owner",
		"kind": map[string]interface{}{
			"key": "value",
		},
	}, data)
}

// Test_UnmarshalYAML tests the UnmarshalYAML function.
func Test_UnmarshalYAML(t *testing.T) {
	t.Parallel()
	r := &Resource{}
	err := r.UnmarshalYAML(func(v interface{}) error {
		*v.(*map[string]interface{}) = map[string]interface{}{
			"name":        "name",
			"description": "description",
			"owner":       "owner",
			"kind": map[string]interface{}{
				"key": "value",
			},
		}
		return nil
	})
	assert.NoError(t, err)
	assert.Equal(t, &Resource{
		Kind:        "kind",
		Name:        "name",
		Description: "description",
		Owner:       "owner",
		Data: map[string]interface{}{
			"key": "value",
		},
	}, r)
}

// Test_UnmarshalYAML_InvalidKind tests the UnmarshalYAML function with an invalid kind.
func Test_UnmarshalYAML_InvalidKind(t *testing.T) {
	t.Parallel()
	r := &Resource{}
	err := r.UnmarshalYAML(func(v interface{}) error {
		*v.(*map[string]interface{}) = map[string]interface{}{
			"name":        "name",
			"description": "description",
			"owner":       "owner",
		}
		return nil
	})
	assert.Error(t, err)
}

// Test_UnmarshalYAML_MultipleKind tests the UnmarshalYAML function with multiple kinds.
func Test_UnmarshalYAML_MultipleKind(t *testing.T) {
	t.Parallel()
	r := &Resource{}
	err := r.UnmarshalYAML(func(v interface{}) error {
		*v.(*map[string]interface{}) = map[string]interface{}{
			"name":        "name",
			"description": "description",
			"owner":       "owner",
			"kind":        "value",
			"other":       "value",
		}
		return nil
	})
	assert.Error(t, err)
}

// Test_UnmarshalYAML_unmarshalError tests the UnmarshalYAML function with an unmarshal error.
func Test_UnmarshalYAML_unmarshalError(t *testing.T) {
	t.Parallel()
	r := &Resource{}
	err := r.UnmarshalYAML(func(v interface{}) error {
		return fmt.Errorf("error")
	})
	assert.Error(t, err)
}

// Test_UnmarshalYAML_DuplicateKey tests the UnmarshalYAML function with a duplicate key.
func Test_UnmarshalYAML_DuplicateKey(t *testing.T) {
	t.Parallel()
	r := &Resource{}
	badYaml := []byte(`
name: name
description: description
owner: owner
kind: value
kind: value
`)
	err := r.UnmarshalYAML(func(v interface{}) error {
		return yaml.Unmarshal(badYaml, v)
	})
	assert.Error(t, err)
}

// Test_MarshalJSON tests the MarshalJSON function.
func Test_MarshalJSON(t *testing.T) {
	t.Parallel()
	r, err := New("kind", "name", "description", "owner")
	assert.NoError(t, err)
	r.Data = map[string]interface{}{
		"key": "value",
	}
	data, err := r.MarshalJSON()
	assert.NoError(t, err)
	jdata := map[string]interface{}{}
	err = json.Unmarshal(data, &jdata)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"name":        "name",
		"description": "description",
		"owner":       "owner",
		"kind": map[string]interface{}{
			"key": "value",
		},
	}, jdata)
}

// Test_UnmarshalJSON tests the UnmarshalJSON function.
func Test_UnmarshalJSON(t *testing.T) {
	t.Parallel()
	r := &Resource{}
	err := r.UnmarshalJSON([]byte(`
{
	"name": "name",
	"description": "description",
	"owner": "owner",
	"kind": {
		"key": "value"
	}
}
`))
	assert.NoError(t, err)
	assert.Equal(t, &Resource{
		Kind:        "kind",
		Name:        "name",
		Description: "description",
		Owner:       "owner",
		Data: map[string]interface{}{
			"key": "value",
		},
	}, r)
}

// Test_UnmarshalJSON_InvalidKind tests the UnmarshalJSON function with an invalid kind.
func Test_UnmarshalJSON_InvalidKind(t *testing.T) {
	t.Parallel()
	r := &Resource{}
	err := r.UnmarshalJSON([]byte(`
{
	"name": "name",
	"description": "description",
	"owner": "owner"
}
`))
	assert.Error(t, err)
}

// Test_UnmarshalJSON_MultipleKind tests the UnmarshalJSON function with multiple kinds.
func Test_UnmarshalJSON_MultipleKind(t *testing.T) {
	t.Parallel()
	r := &Resource{}
	err := r.UnmarshalJSON([]byte(`
{
	"name": "name",
	"description": "description",
	"owner": "owner",
	"kind": "value",
	"other": "value"
}
`))
	assert.Error(t, err)
}

// Test_UnmarshalJSON_unmarshalError tests the UnmarshalJSON function with an unmarshal error.
func Test_UnmarshalJSON_unmarshalError(t *testing.T) {
	t.Parallel()
	r := &Resource{}
	err := r.UnmarshalJSON([]byte(`error`))
	assert.Error(t, err)
}

// Test_UnmarshalJSON_DuplicateKey tests the UnmarshalJSON function with a duplicate key.
func Test_UnmarshalJSON_DuplicateKey(t *testing.T) {
	t.Parallel()
	r := &Resource{}
	badJSON := []byte(`
{
	"name": "name",
	"description": "description",
	"owner": "owner",
	"kind": "value",
	"kind": "value"
}
`)
	err := r.UnmarshalJSON(badJSON)
	assert.NoError(t, err)
	assert.Equal(t, &Resource{
		Kind:        "kind",
		Name:        "name",
		Description: "description",
		Owner:       "owner",
		Data:        "value",
	}, r)

}
