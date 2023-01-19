package usecases

import (
	"github.com/ervitis/sink-float/master/config"
	"github.com/ervitis/sink-float/master/domain"
	"sync"
	"time"
)

type attackClient struct {
	battleships domain.BattleShips

	clients map[string]config.SinkClient
}

func (a *attackClient) Attack() {
	tik := time.NewTicker(30 * time.Second)

	wg := sync.WaitGroup{}
	wg.Add(len(a.battleships))

	go func() {
		for {
			select {
			case <-tik.C:
				for _, battleship := range a.battleships {
					battleship.LaunchMissile()
				}
			}
		}
	}()

	wg.Wait()
}

func (a *attackClient) CheckShips() {

}

type AttackClient interface {
	Attack()
	CheckShips()
}

func NewAttackClient(battleships domain.BattleShips) AttackClient {
	clients := config.LoadClients()

	data := make(config.AddressesSinkClient)
	for _, client := range clients {
		data[client.URI()] = client
	}
	return &attackClient{
		battleships: battleships,
		clients:     data,
	}
}
