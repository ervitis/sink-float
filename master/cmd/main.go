package main

import (
	"context"
	"fmt"
	"github.com/ervitis/sink-float/master/config"
	grpcHandler "github.com/ervitis/sink-float/master/handlers/grpc"
	"github.com/ervitis/sink-float/master/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os/signal"
	"sync"
	"syscall"
)

func init() {
	config.LoadAppConfig()
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	ls, err := net.Listen("tcp", fmt.Sprintf(":%d", config.App.Server.Port))
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	handler := registry.HandleGRPCHandlerProvider(config.App)
	grpcHandler.RegisterSinkFleetServiceServer(server, handler.SinkHandler)
	reflection.Register(server)

	go func() {
		<-ctx.Done()
		server.GracefulStop()
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		log.Panicln(server.Serve(ls))
	}()

	wg.Wait()
}
