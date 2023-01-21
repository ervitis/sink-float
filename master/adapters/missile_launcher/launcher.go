package missile_launcher

import (
	"context"
	"errors"
	"github.com/ervitis/sink-float/master/config"
	"github.com/golang/protobuf/ptypes/empty"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpc2 "github.com/ervitis/sink-float/master/adapters/grpc_impl"
	"github.com/ervitis/sink-float/master/domain"
)

type MissileLauncher grpc2.MasterSinkFleetServiceClient
type dataLauncher struct {
	launcher MissileLauncher
	client   config.SinkClient
}

type launcher struct {
	data []dataLauncher
}

type Launcher interface {
	Launch() error
	Check() error
}

var NoHit = errors.New("NO HIT")
var NoAlive = errors.New("NO ALIVE")

func New(addressesClients config.AddressesSinkClient) Launcher {
	l := new(launcher)
	for uri, client := range addressesClients {
		cn, err := grpc.Dial(
			uri,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			panic(err)
		}
		dl := dataLauncher{
			launcher: grpc2.NewMasterSinkFleetServiceClient(cn),
			client:   client,
		}
		l.data = append(l.data, dl)
	}

	return l
}

func (l launcher) Launch() error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	n := domain.GenerateRandNumber(1, 5)

	// TODO fix this point
	r, err := l.client.Attack(ctx, &grpc2.AtomicMissile{
		GuessNumber: uint32(n),
	})
	if err != nil {
		return err
	}

	if r.GetHit() {
		return nil
	}
	return NoHit
}

func (l launcher) Check() error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	r, err := l.data.CheckIsAlive(ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	if !r.GetIsAlive() {
		return NoAlive
	}
	return nil
}
