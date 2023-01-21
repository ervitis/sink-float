package main

import (
	"context"
	"fmt"
	"github.com/ervitis/sink-float/master/adapters/grpc_impl"
	"github.com/ervitis/sink-float/master/adapters/handlers/grpc"
	"github.com/ervitis/sink-float/master/registry"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ervitis/sink-float/master/config"
	"github.com/ervitis/sink-float/master/domain"
)

func init() {
	config.LoadAppConfig()
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL)
	defer stop()

	grpcSrvSvc := grpc.New()
	handler := registry.HandleGRPCHandlerProvider(config.App)
	grpc_impl.RegisterMasterSinkFleetServiceServer(grpcSrvSvc.Server(), handler.SinkHandler)

	game := domain.New()
	fmt.Println(game)

	go func() {
		<-ctx.Done()
		log.Println("Stopping")
		grpcSrvSvc.Shutdown()
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		log.Printf("Serving grpc server on %d\n", config.App.Server.Port)
		if err := grpcSrvSvc.Serve(); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
	os.Exit(0)
}
