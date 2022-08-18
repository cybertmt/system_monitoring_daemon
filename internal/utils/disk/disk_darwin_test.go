//go:build darwin
// +build darwin

package disk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStat(t *testing.T) {
	t.Run("test success get stats", func(t *testing.T) {
		diskStat, err := Get()

		require.NoError(t, err)
		require.NotNil(t, diskStat.KBt)
		require.IsType(t, 1.0, diskStat.KBt)
		require.NotNil(t, diskStat.MBs)
		require.IsType(t, 1.0, diskStat.MBs)
		require.NotNil(t, diskStat.TPS)
		require.IsType(t, 1, diskStat.TPS)
	})
}
