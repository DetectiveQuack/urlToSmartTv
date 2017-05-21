package main

import (
	"bufio"
	"device"
	"fmt"
	"os"

	"github.com/huin/goupnp"
)

const (
	searchTarget = "urn:schemas-upnp-org:device:MediaRenderer:1"
)

var (
	rootDevices    []goupnp.MaybeRootDevice
	selectedDevice goupnp.MaybeRootDevice
)

//'ConnectionManager', 'GetProtocolInfo'
func main() {

	devices, err := searchDevices()

	if err != nil || len(devices) == 0 {
		fmt.Println("Cannot find any devices:", err)
		return
	}

	// Put devices select into another file/package handling multiple devices
	if len(devices) == 1 {
		selectedDevice = devices[0]
	}
	fmt.Print("Enter Video url: ")
	reader := bufio.NewReader(os.Stdin)
	uri, _ := reader.ReadString('\n')
	device.Play(selectedDevice, uri)
}

func searchDevices() (rootDevices []goupnp.MaybeRootDevice, err error) {
	devices, err := goupnp.DiscoverDevices(searchTarget)
	rootDevices = devices

	return devices, err
}
