package git

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/neuralnorthwest/tpology/git/mock_git"
	"github.com/stretchr/testify/assert"
)

// newTestRepo instantiates a test repository.
func newTestRepo(t *testing.T, c *Cache) *Repository {
	t.Helper()
	return c.New("https://github.com/github/hub.git", "master")
}

// execLog calls exec and sends the output to t.Log.
func execLog(t *testing.T, name string, arg ...string) error {
	t.Helper()
	cmd := exec.Command(name, arg...)
	output, err := cmd.CombinedOutput()
	t.Log(string(output))
	return err
}

// mustExecLog calls exec and sends the output to t.Log. If the command fails,
// it calls t.Fatal.
func mustExecLog(t *testing.T, name string, arg ...string) {
	t.Helper()
	err := execLog(t, name, arg...)
	if err != nil {
		t.Fatal(err)
	}
}

// newLocalTestRepo instantiates a test repository with a local path.
func newLocalTestRepo(t *testing.T, c *Cache) (*Repository, func()) {
	tmprepo, err := os.MkdirTemp("", "git-test-")
	t.Logf("tmprepo: %s", tmprepo)
	assert.NoError(t, err)
	mustClean := true
	mc := &mustClean
	defer func() {
		if *mc {
			t.Logf("cleaning up %s", tmprepo)
			os.RemoveAll(tmprepo)
		}
	}()
	mustExecLog(t, "git", "-C", tmprepo, "init")
	mustExecLog(t, "git", "-C", tmprepo, "checkout", "-b", "main")
	mustExecLog(t, "git", "-C", tmprepo, "config", "user.name", "Test User")
	mustExecLog(t, "git", "-C", tmprepo, "config", "user.email", "test@user")
	mustExecLog(t, "git", "-C", tmprepo, "commit", "--allow-empty", "-m", "Initial commit")
	*mc = false
	return c.New(tmprepo, "main"), func() {
		t.Logf("cleaning up %s", tmprepo)
		err := os.RemoveAll(tmprepo)
		assert.NoError(t, err)
	}
}

// setupCache sets up the Git cache for testing.
func setupCache(t *testing.T, opts ...interface{}) (*Cache, string, func()) {
	t.Helper()
	// create default fs if none is provided
	var f fs = &osFS{}
	if len(opts) > 0 {
		f = opts[0].(fs)
	}
	tmpdir, err := f.MkdirTemp("", "itool-")
	assert.NoError(t, err)
	c := NewCache(tmpdir)
	c.fs = f
	return c, tmpdir, func() {
		err := f.RemoveAll(tmpdir)
		assert.NoError(t, err)
	}
}

// Test_GitCache_NewCache tests the NewCache function.
func Test_GitCache_NewCache(t *testing.T) {
	t.Parallel()
	c, tmpdir, cleanup := setupCache(t)
	defer cleanup()
	assert.Equal(t, c.CachePath, tmpdir)
}

// Test_cleanURL tests the cleanURL function.
func Test_cleanURL(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "https/github.com/github/hub.git", cleanURL("https://github.com/github/hub.git"))
	assert.Equal(t, "https/user/github.com/github/hub.git", cleanURL("https://user@github.com/github/hub.git"))
	assert.Equal(t, "https/user/github.com/github/hub.git", cleanURL("https://user:pass@github.com/github/hub.git"))
	assert.Equal(t, "git/github.com/github/hub.git", cleanURL("git@github.com:github/hub.git"))
	assert.Equal(t, "ssh/github.com/github/hub.git", cleanURL("ssh://github.com/github/hub.git"))
}

// Test_GitCache_New tests the New function.
func Test_GitCache_New(t *testing.T) {
	t.Parallel()
	c, tmpdir, cleanup := setupCache(t)
	defer cleanup()
	repo := c.New("https://github.com/github/hub.git", "master")
	assert.Equal(t, "https://github.com/github/hub.git", repo.URL)
	assert.Equal(t, "master", repo.MainBranch)
	subPath := filepath.Join("https", "github.com", "github", "hub.git")
	assert.Equal(t, filepath.Join(tmpdir, subPath), repo.Dir)
}

// Test_GitRepository_IsCloned tests the IsCloned function.
func Test_GitRepository_IsCloned(t *testing.T) {
	t.Parallel()
	c, _, cleanup := setupCache(t)
	defer cleanup()
	repo, cleanup := newLocalTestRepo(t, c)
	defer cleanup()
	assert.False(t, repo.IsCloned())
	// Now clone
	assert.NoError(t, repo.Clone())
	assert.True(t, repo.IsCloned())
}

// Test_GitRepository_Clone_Local tests the Clone function.
func Test_GitRepository_Clone(t *testing.T) {
	t.Parallel()
	c, _, cleanup := setupCache(t)
	defer cleanup()
	repo, cleanup := newLocalTestRepo(t, c)
	defer cleanup()
	assert.NoError(t, repo.Clone())
	// Now clone again
	assert.NoError(t, repo.Clone())
}

// Test_GitRepository_Clone_Remote tests the Clone function.
func Test_GitRepository_Clone_Remote(t *testing.T) {
	t.Parallel()
	c, _, cleanup := setupCache(t)
	defer cleanup()
	repo := newTestRepo(t, c)
	assert.NoError(t, repo.Clone())
	// Now clone again
	assert.NoError(t, repo.Clone())
}

// Test_GitRepository_Clone_MkdirFails tests the Clone function when the
// directory cannot be created.
func Test_GitRepository_Clone_MkdirFails(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	mockfs := mock_git.NewMockfs(mockCtrl)
	mockfs.EXPECT().Stat(gomock.Any()).Return(nil, errors.New("stat failed"))
	mockfs.EXPECT().MkdirAll(gomock.Any(), gomock.Any()).Return(errors.New("mkdirall failed"))

	c, _, cleanup := setupCache(t)
	defer cleanup()
	c.fs = mockfs
	repo, cleanup := newLocalTestRepo(t, c)
	defer cleanup()
	err := repo.Clone()
	assert.Error(t, err)
}

// Test_GitRepository_Exec tests the Exec function.
func Test_GitRepository_Exec(t *testing.T) {
	t.Parallel()
	c, _, cleanup := setupCache(t)
	defer cleanup()
	repo, cleanup := newLocalTestRepo(t, c)
	defer cleanup()
	assert.NoError(t, repo.Clone())
	assert.NoError(t, repo.Exec("status"))
	// now run a non-existent command
	assert.Error(t, repo.Exec("non-existent-command"))
}

// Test_GitRepository_Remove tests the Remove function.
func Test_GitRepository_Remove(t *testing.T) {
	t.Parallel()
	c, _, cleanup := setupCache(t)
	defer cleanup()
	repo, cleanup := newLocalTestRepo(t, c)
	defer cleanup()
	assert.NoError(t, repo.Clone())
	assert.True(t, repo.IsCloned())
	assert.NoError(t, repo.Remove())
	assert.False(t, repo.IsCloned())
	_, err := os.Stat(repo.Dir)
	assert.Error(t, err)
}

// Test_GitRepository_Clean tests the Clean function.
func Test_GitRepository_Clean(t *testing.T) {
	t.Parallel()
	c, _, cleanup := setupCache(t)
	defer cleanup()
	repo, cleanup := newLocalTestRepo(t, c)
	defer cleanup()
	assert.NoError(t, repo.Clone())
	// Create a file in the repo
	f, err := os.Create(filepath.Join(repo.Dir, "test.txt"))
	assert.NoError(t, err)
	f.Close()
	assert.NoError(t, repo.Clean())
	// Now check the file is gone
	_, err = os.Stat(filepath.Join(repo.Dir, "test.txt"))
	assert.Error(t, err)
}
