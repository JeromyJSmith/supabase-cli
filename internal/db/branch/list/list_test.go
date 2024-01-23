package list

import (
	"bytes"
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/supabase/cli/internal/utils"
)

func TestListCommand(t *testing.T) {
	t.Run("lists all branches", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fsys, utils.CurrBranchPath, []byte("main"), 0644))
		base := filepath.Dir(utils.CurrBranchPath)
		require.NoError(t, fsys.Mkdir(filepath.Join(base, "main"), 0755))
		require.NoError(t, fsys.Mkdir(filepath.Join(base, "test"), 0755))
		// Run test
		var out bytes.Buffer
		require.NoError(t, Run(fsys, &out))
		// Validate output
		lines := strings.Split(out.String(), "\n")
		assert.ElementsMatch(t, []string{
			"* main",
			"  test",
			"",
		}, lines)
	})

	t.Run("lists without current branch", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		base := filepath.Dir(utils.CurrBranchPath)
		require.NoError(t, fsys.Mkdir(filepath.Join(base, "main"), 0755))
		require.NoError(t, fsys.Mkdir(filepath.Join(base, "test"), 0755))
		// Run test
		var out bytes.Buffer
		require.NoError(t, Run(fsys, &out))
		// Validate output
		lines := strings.Split(out.String(), "\n")
		assert.ElementsMatch(t, []string{
			"  main",
			"  test",
			"",
		}, lines)
	})

	t.Run("lists uninitialized branch", func(t *testing.T) {
		require.NoError(t, Run(afero.NewMemMapFs(), io.Discard))
	})

	t.Run("throws error on unreadable directory", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		_, err := fsys.Create(filepath.Dir(utils.CurrBranchPath))
		require.NoError(t, err)
		// Run test
		require.Error(t, Run(fsys, io.Discard))
	})

	t.Run("correctly lists branches when current branch file is missing", func(t *testing.T) {
		// Setup in-memory fs without current branch
		fsys := afero.NewMemMapFs()
		base := filepath.Dir(utils.CurrBranchPath)
		require.NoError(t, fsys.Mkdir(filepath.Join(base, "feature"), 0755))
		// Run test for missing current branch file
		var out bytes.Buffer
		require.NoError(t, Run(fsys, &out))
		// Validate output
		lines := strings.Split(out.String(), "\n")
		assert.ElementsMatch(t, []string{
			"  feature",
			"",
		}, lines)
	})

	t.Run("fails gracefully when branch directory is missing", func(t *testing.T) {
		// Setup in-memory fs with current branch but no branch directory
		fsys := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fsys, utils.CurrBranchPath, []byte("feature"), 0644))
		// Run test for missing branch directory
		err := Run(fsys, io.Discard)
		require.Error(t, err)
		assert.EqualError(t, err, "branches directory not found")
	})
}
