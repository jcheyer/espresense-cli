package network

import (
	"fmt"

	"github.com/jcheyer/espresense-cli/internal/espresense"
)

type Network struct {
	baseStations []*espresense.Instance
}

func New() *Network {
	return &Network{
		baseStations: make([]*espresense.Instance, 0),
	}
}

func (n *Network) AddBaseStation(ip string) error {
	s, err := espresense.New(ip, nil)
	if err != nil {
		return err
	}

	n.baseStations = append(n.baseStations, s)

	return nil
}

func (n *Network) DeviceByID(ID string) []*espresense.Device {
	devices := make([]*espresense.Device, 0)
	for _, station := range n.baseStations {
		d, err := station.DeviceByID(ID)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		devices = append(devices, d)
	}
	return devices
}
