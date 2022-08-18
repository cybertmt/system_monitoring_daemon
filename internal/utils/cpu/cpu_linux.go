//go:build linux
// +build linux

package cpu

import (
	"github.com/cybertmt/system_monitoring_daemon/internal/app"
	"github.com/cybertmt/system_monitoring_daemon/internal/utils"
)

func Get() (*app.CPUStats, error) {
	return nil, utils.ErrNotImplemented
}
