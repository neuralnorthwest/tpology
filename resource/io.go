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
	return Load(inf)
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
