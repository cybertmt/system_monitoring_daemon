package pipeline

import (
	"log"

	"github.com/cybertmt/system_monitoring_daemon/internal/app"
	"github.com/cybertmt/system_monitoring_daemon/internal/config"
	"github.com/cybertmt/system_monitoring_daemon/internal/utils/cpu"
	"github.com/cybertmt/system_monitoring_daemon/internal/utils/disk"
	"github.com/cybertmt/system_monitoring_daemon/internal/utils/load"
)

func GetStages(statsConfig config.StatsConfig) []Stage {
	var stages []Stage

	if statsConfig.LoadAvg {
		stages = append(stages, stageGenerator(
			"Load Average Stage",
			func(stat app.SystemStats) app.SystemStats {
				loadAvg, err := load.Get()
				if err != nil {
					log.Printf("failed to get load average: %s", err)
				}
				stat.Load = loadAvg

				return stat
			}),
		)
	}

	if statsConfig.CPU {
		stages = append(stages, stageGenerator(
			"CPU stage",
			func(stat app.SystemStats) app.SystemStats {
				cpuStat, err := cpu.Get()
				if err != nil {
					log.Printf("failed to get CPU usage: %s", err)
				}
				stat.CPU = cpuStat

				return stat
			}),
		)
	}

	if statsConfig.Disk {
		stages = append(stages, stageGenerator(
			"Disk stage",
			func(stat app.SystemStats) app.SystemStats {
				diskStats, err := disk.Get()
				if err != nil {
					log.Printf("failed to get disk stats: %s", err)
				}
				stat.Disk = diskStats

				return stat
			}),
		)
	}

	return stages
}

func stageGenerator(_ string, f func(s app.SystemStats) app.SystemStats) Stage {
	return func(in In) Out {
		out := make(Bi)
		go func() {
			defer close(out)
			for v := range in {
				out <- f(v.(app.SystemStats))
			}
		}()
		return out
	}
}
