package espresense

import (
	"errors"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestDevices(t *testing.T) {
	r := require.New(t)
	defer gock.Off()
	gock.Observe(gock.DumpRequest)

	testIP := "192.168.1.1"
	testClient := &http.Client{}

	gock.InterceptClient(testClient)

	// Init
	gock.New("http://" + testIP).
		Get("/json/devices").
		Reply(200).
		File("testdata/devices.json")

	i, err := New(testIP, testClient)
	r.NotNil(i)
	r.NoError(err)
	r.Equal("thisroom", i.name)

	gock.New("http://" + testIP).
		Get("/json/devices").
		ReplyError(errors.New("bla"))

	d, err := i.Devices()
	r.Error(err)
	r.Nil(d)

	gock.New("http://" + testIP).
		Get("/json/devices").
		Reply(200).
		File("testdata/devices.json")

	d, err = i.Devices()
	r.NoError(err)
	r.NotNil(d)
	r.Len(d, 3)

	r.True(gock.IsDone())
}

func TestDeviceByID(t *testing.T) {

	r := require.New(t)
	defer gock.Off()
	gock.Observe(gock.DumpRequest)

	testIP := "192.168.1.1"
	testClient := &http.Client{}

	gock.InterceptClient(testClient)

	gock.New("http://" + testIP).
		Get("/devices").
		Persist().
		Reply(200).
		File("testdata/devices.json")

	i, err := New(testIP, testClient)
	r.NoError(err)

	d, err := i.DeviceByID("fasd")
	r.Error(err)
	r.Nil(d)

	testID := "iBeacon:e5ca1ade-f007-ba11-0000-000000000000-156-31179"

	d, err = i.DeviceByID(testID)
	r.NoError(err)
	r.Equal(testID, d.ID)

}

func TestDeviceByName(t *testing.T) {

	r := require.New(t)
	defer gock.Off()
	gock.Observe(gock.DumpRequest)

	testIP := "192.168.1.1"
	testClient := &http.Client{}

	gock.InterceptClient(testClient)

	gock.New("http://" + testIP).
		Get("/devices").
		Persist().
		Reply(200).
		File("testdata/devices.json")

	i, err := New(testIP, testClient)
	r.NoError(err)

	d, err := i.DeviceByName("fasd")
	r.Error(err)
	r.Nil(d)

	testName := "otherroom"

	d, err = i.DeviceByName(testName)
	r.NoError(err)
	r.Equal(testName, d.Name)

}
