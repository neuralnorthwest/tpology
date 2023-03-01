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
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
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
	// PreHook is a function that is called before a Git command is executed.
	PreHook func(*exec.Cmd) error
	// PostHook is a function that is called after a Git command is executed.
	PostHook func(*exec.Cmd, error) error
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

// CloneIfNotCloned clones the Git repository if it is not cloned.
func (r *Repository) CloneIfNotCloned() error {
	if !r.IsCloned() {
		return r.Clone()
	}
	return nil
}

// Clone clones the Git repository.
func (r *Repository) Clone(args ...string) error {
	if !r.IsCloned() {
		if err := r.fs.MkdirAll(filepath.Dir(r.Dir), 0755); err != nil {
			return err
		}
		var cmd *exec.Cmd
		if r.MainBranch == "" {
			cmd = exec.Command("git", append(append([]string{"clone"}, args...), r.URL, r.Dir)...)
		} else {
			cmd = exec.Command("git", append(append([]string{"clone", "-b", r.MainBranch}, args...), r.URL, r.Dir)...)
		}
		if r.PreHook != nil {
			if herr := r.PreHook(cmd); herr != nil {
				return herr
			}
		}
		err := cmd.Run()
		if r.PostHook != nil {
			if herr := r.PostHook(cmd, err); herr != nil {
				return herr
			}
		}
		return err
	} else {
		return fmt.Errorf("repository already cloned: %s", r.URL)
	}
}

// Remove removes the Git repository.
func (r *Repository) Remove() error {
	return r.fs.RemoveAll(r.Dir)
}

// Clean cleans the Git repository.
func (r *Repository) Clean() error {
	return r.Exec("clean", "-fdx")
}

// Checkout checks out the specified branch of the Git repository.
func (r *Repository) Checkout(branch string) error {
	return r.Exec("checkout", branch)
}

// IsClean returns true if the Git repository is clean.
func (r *Repository) IsClean() bool {
	out, err := r.ExecOutput("status", "--porcelain")
	return err == nil && out == ""
}

// Lock locks the Git repository. It returns an unlock function and an error.
func (r *Repository) Lock() (func(), error) {
	// Exclusive create the lock file and write our PID to it
	lockFile := filepath.Join(r.Dir, ".lock")
	if f, err := r.fs.OpenFile(lockFile, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644); err == nil {
		defer f.Close()
		if _, err := f.WriteString(strconv.Itoa(os.Getpid())); err != nil {
			return nil, err
		}
		return func() {
			_ = r.fs.Remove(lockFile)
		}, nil
	} else {
		otherPID := 0
		if f, err := r.fs.Open(lockFile); err == nil {
			if _, err := fmt.Fscanf(f, "%d", &otherPID); err == nil {
				f.Close()
			}
		}
		return nil, fmt.Errorf("repository already locked by PID %d: %s", otherPID, r.URL)
	}
}

// LockerPID returns the PID of the process that locked the Git repository.
func (r *Repository) LockerPID() (int, error) {
	lockFile := filepath.Join(r.Dir, ".lock")
	if f, err := r.fs.Open(lockFile); err == nil {
		defer f.Close()
		var pid int
		if _, err := fmt.Fscanf(f, "%d", &pid); err == nil {
			return pid, nil
		}
	}
	return 0, fmt.Errorf("repository not locked: %s", r.URL)
}

// Exec executes a command in the Git repository.
func (r *Repository) Exec(args ...string) error {
	args = append([]string{"-C", r.Dir}, args...)
	cmd := exec.Command("git", args...)
	if r.PreHook != nil {
		if herr := r.PreHook(cmd); herr != nil {
			return herr
		}
	}
	err := cmd.Run()
	if r.PostHook != nil {
		if herr := r.PostHook(cmd, err); herr != nil {
			return herr
		}
	}
	return err
}

// ExecOutput executes a command in the Git repository and returns its output.
func (r *Repository) ExecOutput(args ...string) (string, error) {
	args = append([]string{"-C", r.Dir}, args...)
	cmd := exec.Command("git", args...)
	if r.PreHook != nil {
		if herr := r.PreHook(cmd); herr != nil {
			return "", herr
		}
	}
	out, err := cmd.Output()
	if r.PostHook != nil {
		if herr := r.PostHook(cmd, err); herr != nil {
			return "", herr
		}
	}
	return string(out), err
}
