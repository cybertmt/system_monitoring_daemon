//go:build linux
// +build linux

package disk

import (
	"github.com/cybertmt/system_monitoring_daemon/internal/app"
	"github.com/cybertmt/system_monitoring_daemon/internal/utils"
)

func Get() (*app.DiskStats, error) {
	return nil, utils.ErrNotImplemented
}
