package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type SinkClient struct {
	Address string
	Port    int
	Alive   bool
}

type AddressesSinkClient map[string]SinkClient

func (s SinkClient) URI() string {
	return fmt.Sprintf("%s:%d", s.Address, s.Port)
}

type SinkClients []SinkClient

func LoadClients() SinkClients {
	path, err := filepath.Abs("clients")
	if err != nil {
		log.Println(err)
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil
	}

	defer func() {
		_ = f.Close()
	}()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)

	clients := make(SinkClients, 0)
	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), "#") {
			continue
		}
		line := strings.Split(sc.Text(), ":")
		if len(line) != 2 {
			continue
		}
		port, err := strconv.Atoi(line[1])
		if err != nil {
			continue
		}

		clients = append(clients, SinkClient{
			Address: line[0],
			Port:    port,
			Alive:   true,
		})
	}
	return clients
}
