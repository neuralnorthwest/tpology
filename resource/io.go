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
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// LoadFile loads a manifest.
func LoadFile(path string) ([]*Resource, error) {
	inf, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer inf.Close()
	resources, err := Load(inf)
	if err != nil {
		return nil, err
	}
	for _, r := range resources {
		r.loadedFrom = path
	}
	return resources, nil
}

// Load loads a manifest from a reader.
func Load(r io.Reader) ([]*Resource, error) {
	resources := []*Resource{}
	dec := yaml.NewDecoder(r)
	for {
		r := &Resource{}
		if err := dec.Decode(r); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		resources = append(resources, r)
	}
	return resources, nil
}
