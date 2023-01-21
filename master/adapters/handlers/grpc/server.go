package grpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/ervitis/sink-float/master/adapters/grpc_impl"
	"net"

	grpcLib "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ervitis/sink-float/master/config"
	"github.com/ervitis/sink-float/master/usecases"
)

type SinkFleetHandler = grpc_impl.MasterSinkFleetServiceServer

type sinkFleetHandler struct {
	usecases.SinkUseCase
}

type sinkGrpcServer struct {
	conn net.Listener
	srv  *grpcLib.Server
}

func (ss sinkGrpcServer) Server() *grpcLib.Server {
	return ss.srv
}

func (ss sinkGrpcServer) Serve() error {
	if err := ss.srv.Serve(ss.conn); err != nil && !errors.Is(err, grpcLib.ErrServerStopped) {
		return err
	}
	return nil
}

func (ss sinkGrpcServer) Shutdown() {
	ss.srv.GracefulStop()
}

type SinkGrpcOperations interface {
	Serve() error
	Shutdown()
	Server() *grpcLib.Server
}

func (s sinkFleetHandler) Attack(ctx context.Context, request *grpc_impl.AtomicMissile) (*grpc_impl.MissileResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewSinkFleetHandler(svc usecases.SinkUseCase) SinkFleetHandler {
	return &sinkFleetHandler{svc}
}

func New() SinkGrpcOperations {
	ls, err := net.Listen("tcp", fmt.Sprintf(":%d", config.App.Server.Port))
	if err != nil {
		panic(err)
	}
	server := grpcLib.NewServer()
	reflection.Register(server)
	return &sinkGrpcServer{
		conn: ls,
		srv:  server,
	}
}
