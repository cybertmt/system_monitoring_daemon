//go:build linux
// +build linux

package load

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/cybertmt/system_monitoring_daemon/internal/app"
)

func Get() (*app.LoadStats, error) {
	stat, err := fileAvg()
	if err != nil {
		stat, err = sysInfoAvg()
	}
	return stat, err
}

func sysInfoAvg() (*app.LoadStats, error) {
	var info syscall.Sysinfo_t
	err := syscall.Sysinfo(&info)
	if err != nil {
		return nil, err
	}

	return &app.LoadStats{
		Load1:  float64(info.Loads[0]) / float64(1<<16),
		Load5:  float64(info.Loads[1]) / float64(1<<16),
		Load15: float64(info.Loads[2]) / float64(1<<16),
	}, nil
}

func fileAvg() (*app.LoadStats, error) {
	values, err := readLoadAvgFromFile()
	if err != nil {
		return nil, err
	}

	load1, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		return nil, err
	}

	load5, err := strconv.ParseFloat(values[1], 64)
	if err != nil {
		return nil, err
	}

	load15, err := strconv.ParseFloat(values[2], 64)
	if err != nil {
		return nil, err
	}

	return &app.LoadStats{
		Load1:  load1,
		Load5:  load5,
		Load15: load15,
	}, nil
}

func readLoadAvgFromFile() ([]string, error) {
	procDir := os.Getenv("HOST_PROC")
	if procDir == "" {
		procDir = "/proc"
	}

	line, err := ioutil.ReadFile(filepath.Join(procDir, "/loadavg"))
	if err != nil {
		return nil, err
	}

	return strings.Fields(string(line)), nil
}
