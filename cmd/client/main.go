package main

import (
	"context"
	"errors"
	"flag"
	"io"
	"log"
	"net"

	internalgrpc "github.com/cybertmt/system_monitoring_daemon/internal/server/grpc"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port string
	n, m int
)

func init() {
	flag.StringVar(&port, "port", "50005", "port")
	flag.IntVar(&n, "n", 5, "interval to get statistic (sec)")
	flag.IntVar(&m, "m", 15, "interval to average statistic (sec)")
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(net.JoinHostPort("", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect with server: %s", err)
	}
	defer conn.Close()

	client := internalgrpc.NewSystemStatsStreamServiceClient(conn)

	in := &internalgrpc.RequestMessage{
		Name: uuid.New().String(),
		N:    int32(n),
		M:    int32(m),
	}

	stream, err := client.FetchResponse(context.Background(), in)
	if err != nil {
		log.Printf("open stream error: %s", err)
		return
	}

	log.Print("started fetching")

	done := make(chan bool)
	go func() {
		for {
			resp, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("can not receive: %s", err)
			}
			log.Printf("response received: %s", resp.Title)
		}
	}()

	<-done
	log.Printf("finished receiving")
}
