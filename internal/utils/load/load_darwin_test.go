//go:build darwin
// +build darwin

package load

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStat(t *testing.T) {
	t.Run("test success get stats", func(t *testing.T) {
		loadStat, err := Get()

		require.NoError(t, err)
		require.NotNil(t, loadStat.Load1)
		require.IsType(t, 1.0, loadStat.Load1)
		require.NotNil(t, loadStat.Load5)
		require.IsType(t, 1.0, loadStat.Load5)
		require.NotNil(t, loadStat.Load15)
		require.IsType(t, 1.0, loadStat.Load15)
	})
}
