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

import "os"

// fs lets us mock some filesystem operations for testing.
type fs interface {
	// MkdirTemp creates a temporary directory.
	MkdirTemp(dir, prefix string) (string, error)
	// MkdirAll creates a directory and its parents.
	MkdirAll(path string, perm os.FileMode) error
	// RemoveAll removes a directory and its contents.
	RemoveAll(path string) error
	// Stat returns the FileInfo for a file.
	Stat(name string) (os.FileInfo, error)
	// OpenFile opens a file.
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
	// Remove removes a file.
	Remove(name string) error
	// Open opens a file.
	Open(name string) (*os.File, error)
}

// generate mocks with gomock
//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -package mock_git -source=fs.go -destination mock_git/fs_mock.go . fs

// osFS implements fsMock using the os package.
type osFS struct{}

// MkdirTemp creates a temporary directory.
func (*osFS) MkdirTemp(dir, prefix string) (string, error) {
	return os.MkdirTemp(dir, prefix)
}

// MkdirAll creates a directory and its parents.
func (*osFS) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// RemoveAll removes a directory and its contents.
func (*osFS) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

// Stat returns the FileInfo for a file.
func (*osFS) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

// OpenFile opens a file.
func (*osFS) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

// Remove removes a file.
func (*osFS) Remove(name string) error {
	return os.Remove(name)
}

// Open opens a file.
func (*osFS) Open(name string) (*os.File, error) {
	return os.Open(name)
}
