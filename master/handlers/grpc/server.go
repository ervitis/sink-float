package grpc

import (
	"context"
	"github.com/ervitis/sink-float/master/usecases"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SinkFleetHandler = SinkFleetServiceServer

type sinkFleetHandler struct {
	usecases.SinkUseCase
}

func (s sinkFleetHandler) SendMissile(ctx context.Context, request *MissileRequest) (*MissileResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s sinkFleetHandler) SendKill(ctx context.Context, request *KillRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func NewSinkFleetHandler(svc usecases.SinkUseCase) SinkFleetHandler {
	return &sinkFleetHandler{svc}
}
