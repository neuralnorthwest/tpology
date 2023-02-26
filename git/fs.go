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
}

// generate mocks with gomock
//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -package mock_git -source=fs.go -destination mock_git/fs_mock.go . fs

// osFS implements fsMock using the os package.
type osFS struct{}

// MkdirTemp creates a temporary directory.
func (*osFS) MkdirTemp(dir, prefix string) (string, error) { return os.MkdirTemp(dir, prefix) }

// MkdirAll creates a directory and its parents.
func (*osFS) MkdirAll(path string, perm os.FileMode) error { return os.MkdirAll(path, perm) }

// RemoveAll removes a directory and its contents.
func (*osFS) RemoveAll(path string) error { return os.RemoveAll(path) }

// Stat returns the FileInfo for a file.
func (*osFS) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }
