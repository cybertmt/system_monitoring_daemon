package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"github.com/cybertmt/system_monitoring_daemon/internal/config"
	internalgrpc "github.com/cybertmt/system_monitoring_daemon/internal/server/grpc"
)

var (
	port       string
	configFile string
)

func init() {
	flag.StringVar(&port, "port", "50005", "Daemon port")
	flag.StringVar(&configFile, "config", "configs/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	grpcServer := internalgrpc.NewServer("", port, *cfg)

	go func() {
		if err := grpcServer.Start(); err != nil {
			log.Fatalf("failed to start grpc server: %s", err)
		}
	}()
	defer grpcServer.Stop()

	<-ctx.Done()
	log.Printf("graceful shutting down")
}
