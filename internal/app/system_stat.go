package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type SystemStatsAvg struct {
	Load1  float64
	Load5  float64
	Load15 float64
	User   float64
	System float64
	Idle   float64
	KBt    float64
	TPS    float64
	MBs    float64
}

func (s SystemStatsAvg) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Load average: %.2f %.2f %.2f", s.Load1, s.Load5, s.Load15))
	builder.WriteString(" ")
	builder.WriteString(fmt.Sprintf("CPU usage: %.2f %.2f %.2f", s.User, s.System, s.Idle))
	builder.WriteString(" ")
	builder.WriteString(fmt.Sprintf("Disk: %.2f %.2f %.2f", s.KBt, s.TPS, s.MBs))

	return builder.String()
}

type SystemStats struct {
	ID          uuid.UUID
	CollectedAt time.Time
	Load        *LoadStats
	CPU         *CPUStats
	Disk        *DiskStats
}

func (s SystemStats) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ID: %s ", s.ID))
	builder.WriteString(fmt.Sprintf("Collected at: %s ", s.CollectedAt))
	if s.Load != nil {
		builder.WriteString(fmt.Sprintf("Load average: %s ", s.Load))
	}
	if s.CPU != nil {
		builder.WriteString(fmt.Sprintf("CPU usage: %s ", s.CPU))
	}
	if s.Disk != nil {
		builder.WriteString(fmt.Sprintf("Disk: %s ", s.Disk))
	}

	return builder.String()
}

type LoadStats struct {
	Load1  float64
	Load5  float64
	Load15 float64
}

func (l LoadStats) String() string {
	return fmt.Sprintf("1m: %2f 5m: %2f 15m: %2f", l.Load1, l.Load5, l.Load15)
}

type CPUStats struct {
	User   int
	System int
	Idle   int
}

func (c CPUStats) String() string {
	return fmt.Sprintf("User: %d System: %d Idle: %d", c.User, c.System, c.Idle)
}

type DiskStats struct {
	KBt float64
	TPS int
	MBs float64
}

func (d DiskStats) String() string {
	return fmt.Sprintf("KB/t: %2f tps: %d MB/s: %2f", d.KBt, d.TPS, d.MBs)
}
