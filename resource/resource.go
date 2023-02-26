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

package resource

import (
	"encoding/json"
	"fmt"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// ErrorKindIsReservedWord is the error returned when a resource kind is a reserved word.
	ErrorKindIsReservedWord = Error("resource kind is a reserved word")
)

// Resource is a resource.
type Resource struct {
	// Kind is the kind of the resource.
	Kind string `json:"kind" yaml:"kind"`
	// Name is the name of the resource.
	Name string `json:"name" yaml:"name"`
	// Description is the description of the resource.
	Description string `json:"description" yaml:"description"`
	// Owner is the owner of the resource.
	Owner string `json:"owner" yaml:"owner"`
	// Data is the data of the resource.
	Data interface{} `json:"-" yaml:"-"`
	// loadedFrom is the path to the file the resource was loaded from.
	loadedFrom string
}

// New returns a new resource.
func New(kind, name, description, owner string) (*Resource, error) {
	if KindIsReservedWord(kind) {
		return nil, fmt.Errorf("%w: %s", ErrorKindIsReservedWord, kind)
	}
	return &Resource{
		Kind:        kind,
		Name:        name,
		Description: description,
		Owner:       owner,
	}, nil
}

// MarshalYAML implements the yaml.Marshaler interface.
func (r *Resource) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"name":        r.Name,
		"description": r.Description,
		"owner":       r.Owner,
		r.Kind:        r.Data,
	}, nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (r *Resource) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var data map[string]interface{}
	if err := unmarshal(&data); err != nil {
		return err
	}
	r.Name = getField(data, "name")
	r.Description = getField(data, "description")
	r.Owner = getField(data, "owner")
	delete(data, "name")
	delete(data, "description")
	delete(data, "owner")
	for kind, value := range data {
		// No need to check for reserved words here because all reserved words
		// are already deleted from the data map.
		r.Kind = kind
		r.Data = value
	}
	if len(data) != 1 {
		return fmt.Errorf("resource has more than one kind")
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (r *Resource) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{
		"name":        r.Name,
		"description": r.Description,
		"owner":       r.Owner,
		r.Kind:        r.Data,
	}
	return json.Marshal(data)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (r *Resource) UnmarshalJSON(data []byte) error {
	var dataMap map[string]interface{}
	if err := json.Unmarshal(data, &dataMap); err != nil {
		return err
	}
	r.Name = getField(dataMap, "name")
	r.Description = getField(dataMap, "description")
	r.Owner = getField(dataMap, "owner")
	delete(dataMap, "name")
	delete(dataMap, "description")
	delete(dataMap, "owner")
	for kind, value := range dataMap {
		// No need to check for reserved words here because all reserved words
		// are already deleted from the data map.
		r.Kind = kind
		r.Data = value
	}
	if len(dataMap) != 1 {
		return fmt.Errorf("resource has more than one kind")
	}
	return nil
}

// KindIsReservedWord returns true if the kind is a reserved word.
func KindIsReservedWord(kind string) bool {
	return kind == "name" || kind == "description" || kind == "owner"
}

// getField gets a field from the resource, returning empty string if the field
// is not present.
func getField(data map[string]interface{}, field string) string {
	value, ok := data[field]
	if !ok {
		return ""
	}
	return value.(string)
}
