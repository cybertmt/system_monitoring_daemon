//go:build darwin
// +build darwin

package load

import (
	"unsafe"

	"github.com/cybertmt/system_monitoring_daemon/internal/app"
	"golang.org/x/sys/unix"
)

func Get() (*app.LoadStats, error) {
	type loadavg struct {
		load  [3]uint32
		scale int
	}

	b, err := unix.SysctlRaw("vm.loadavg")
	if err != nil {
		return nil, err
	}
	load := *(*loadavg)(unsafe.Pointer(&b[0]))
	scale := float64(load.scale)

	return &app.LoadStats{
		Load1:  float64(load.load[0]) / scale,
		Load5:  float64(load.load[1]) / scale,
		Load15: float64(load.load[2]) / scale,
	}, nil
}
