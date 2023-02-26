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

package git

import (
	"net/url"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// Cache is the Git cache.
type Cache struct {
	// CachePath is the path to the Git cache directory.
	CachePath string
	// fs is the filesystem interface.
	fs
}

// Repository represents a local clone of a Git repository.
type Repository struct {
	// URL is the URL of the Git repository.
	URL string
	// MainBranch is the name of the main branch of the Git repository.
	MainBranch string
	// Dir is the path to the directory where the Git repository is cloned.
	Dir string
	// fs is the filesystem interface.
	fs
}

// NewCache returns a new Git cache.
func NewCache(cachePath string) *Cache {
	return &Cache{
		CachePath: cachePath,
		fs:        &osFS{},
	}
}

// cleanURL cleans the URL of the Git repository.
func cleanURL(u string) string {
	if parsed, err := url.Parse(u); err == nil {
		// Strip the password, if any
		parsed.User = url.User(parsed.User.Username())
		u = parsed.String()
	}
	u = regexp.MustCompile(`[:@]`).ReplaceAllString(u, "/")
	elems := strings.Split(u, "/")
	cleanElems := []string{}
	for _, elem := range elems {
		if elem == "" {
			continue
		}
		cleanElems = append(cleanElems, strings.ReplaceAll(elem, ":", "/"))
	}
	return strings.Join(cleanElems, "/")
}

// New returns a new Git repository.
func (c *Cache) New(url, mainBranch string) *Repository {
	return &Repository{
		URL:        url,
		MainBranch: mainBranch,
		Dir:        filepath.Join(c.CachePath, cleanURL(url)),
		fs:         c.fs,
	}
}

// isCloned returns true if the Git repository is cloned.
func (r *Repository) IsCloned() bool {
	_, err := r.fs.Stat(r.Dir)
	return err == nil
}

// Exec executes a command in the Git repository.
func (r *Repository) Exec(args ...string) error {
	args = append([]string{"-C", r.Dir}, args...)
	return exec.Command("git", args...).Run()
}

// Clone clones the Git repository.
func (r *Repository) Clone() error {
	if !r.IsCloned() {
		if err := r.fs.MkdirAll(filepath.Dir(r.Dir), 0755); err != nil {
			return err
		}
		return exec.Command("git", "clone", "--depth=1", "--branch", r.MainBranch, r.URL, r.Dir).Run()
	}
	return nil
}

// Remove removes the Git repository.
func (r *Repository) Remove() error {
	return r.fs.RemoveAll(r.Dir)
}

// Clean cleans the Git repository.
func (r *Repository) Clean() error {
	return r.Exec("clean", "-fdx")
}
