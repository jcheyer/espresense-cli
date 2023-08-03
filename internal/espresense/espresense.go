package espresense

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
)

type Instance struct {
	ip     string
	client *resty.Client
	debug  bool
	name   string
}

func New(ip string, client *http.Client) (*Instance, error) {

	if client == nil {
		client = http.DefaultClient
	}

	i := &Instance{
		ip:     ip,
		client: resty.NewWithClient(client),
		debug:  true,
	}

	_, err := i.Name()
	if err != nil {
		return nil, err
	}

	return i, err
}

func (i *Instance) Name() (string, error) {
	if i.name != "" {
		return i.name, nil
	}

	dr, err := i.request()
	if err != nil {
		return "", err
	}
	i.name = dr.Room

	return dr.Room, nil
}

func (i *Instance) request() (*DevicesResponse, error) {
	resp, err := i.client.R().Get(fmt.Sprintf("http://%s/json/devices", i.ip))
	if err != nil {
		return nil, err
	}

	devicesResponse := DevicesResponse{}
	err = json.Unmarshal(resp.Body(), &devicesResponse)
	if err != nil && i.debug {
		spew.Dump(resp.Body())
	}
	return &devicesResponse, nil
}

func (i *Instance) Devices() ([]*Device, error) {
	dr, err := i.request()
	if err != nil {
		return nil, err
	}

	return dr.Devices, nil
}

func (i *Instance) DeviceByID(id string) (*Device, error) {
	d, err := i.Devices()
	if err != nil {
		return nil, err
	}
	for _, device := range d {
		if device.ID == id {
			return device, nil
		}
	}
	return nil, fmt.Errorf("DeviceByID %s not found", id)
}

func (i *Instance) DeviceByName(name string) (*Device, error) {
	d, err := i.Devices()
	if err != nil {
		return nil, err
	}
	for _, device := range d {
		if device.Name == name {
			return device, nil
		}
	}
	return nil, fmt.Errorf("DeviceByName %s not found", name)
}
