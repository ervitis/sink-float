package domain

import (
	"errors"
	"fmt"

	"github.com/ervitis/sink-float/master/adapters/missile_launcher"
)

const ShipSize = 3

type BattleShip struct {
	Size  uint32
	Alive bool
}

type Submarine struct {
	BattleShip

	launcher missile_launcher.Launcher
}

type IBattleShip interface {
	LaunchMissile()
	CheckShipWithRadar() error
}

type BattleShips []IBattleShip

func NewSubmarine() IBattleShip {
	return &Submarine{
		BattleShip{
			Size:  ShipSize,
			Alive: true,
		},
		missile_launcher.New(),
	}
}

func (s Submarine) LaunchMissile() {
	if !s.Alive {
		return
	}

	err := s.launcher.Launch()
	if err != nil && errors.Is(err, missile_launcher.NoHit) {
		fmt.Println("Failed!")
	} else {
		fmt.Println(err)
	}
}

func (s Submarine) CheckShipWithRadar() error {
	if !s.Alive {
		return errors.New("my ship is not alive")
	}
	return s.launcher.Check()
}
