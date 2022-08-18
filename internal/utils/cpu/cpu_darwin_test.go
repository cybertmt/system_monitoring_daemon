//go:build darwin
// +build darwin

package cpu

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStat(t *testing.T) {
	t.Run("test success get stats", func(t *testing.T) {
		cpuStat, err := Get()

		require.NoError(t, err)
		require.NotNil(t, cpuStat.User)
		require.IsType(t, 1, cpuStat.User)
		require.NotNil(t, cpuStat.System)
		require.IsType(t, 1, cpuStat.System)
		require.NotNil(t, cpuStat.Idle)
		require.IsType(t, 1, cpuStat.Idle)
	})
}
