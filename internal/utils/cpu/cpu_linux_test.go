//go:build linux
// +build linux

package cpu

import (
	"github.com/cybertmt/system_monitoring_daemon/internal/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetStat(t *testing.T) {
	t.Run("test not implemented get stats", func(t *testing.T) {
		cpuStat, err := Get()

		require.Nil(t, cpuStat)
		require.ErrorIs(t, err, utils.ErrNotImplemented)
	})
}
