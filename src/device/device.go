package device

import (
	"net/http"
	"net/url"

	"fmt"

	"encoding/xml"

	"io"

	"strings"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/soap"
)

const (
	playServiceID = "urn:upnp-org:serviceId:ConnectionManager"
)

var (
	services   []string
	smartTv    TV
	rootDevice goupnp.MaybeRootDevice
	baseURL    string
	actions    = []string{"ConnectionManager", "RenderingControl", "AVTransport"}
)

// HandleDevice is where all device functionality is, this takes the selected device
func handleDevice(device goupnp.MaybeRootDevice) (tv TV) {
	deviceLocation := device.Location.String()

	response, err := http.Get(deviceLocation)

	if err != nil {
		fmt.Println("Cannot query url, tv may be turned off or not connected to wifi", deviceLocation)
		return
	}

	return decodeDevice(&response.Body)
}

func decodeDevice(xmlResponse *io.ReadCloser) (tv TV) {
	decoder := xml.NewDecoder(*xmlResponse)

	decoder.Decode(&tv)

	fmt.Println("Connected to:", tv.DeviceInfo.Name)

	smartTv = tv

	return tv
}

// Play url on selected device
func Play(device goupnp.MaybeRootDevice, uri string) {
	baseURL = "http://" + device.Location.Host

	handleDevice(device)
	prepareForConnection()
	setAvTransportURI(uri)

	in := PlayInArgs{"0", "1"}
	out := OutAction{}
	err := sendSoapRequest(actions[2], "Play", &in, &out)

	fmt.Println(out, err)
}

func getService(action string) Service {
	service := Service{}

	for _, element := range smartTv.DeviceInfo.Services {
		if strings.Contains(element.ServiceID, action) {
			service = element
			break
		}
	}

	return service
}

func prepareForConnection() {
	in := PrepareForConnectionInArgs{
		"http-get:*video/mp4:*",
		"",
		"-1",
		"Input",
	}
	out := PrepareForConnectionOutArgs{}

	err := sendSoapRequest(actions[0], "PrepareForConnection", &in, &out)

	fmt.Println(out, err)
}

func setAvTransportURI(uri string) {
	in := SetAVTransportURIInArgs{"0", uri, ""}
	out := SetAVTransportURIOutArgs{}
	fmt.Println(uri)
	err := sendSoapRequest(actions[2], "SetAVTransportURI", &in, &out)

	fmt.Println(out, err)
}

func sendSoapRequest(action string, actionName string, in interface{}, out interface{}) (err error) {
	service := getService(action)

	parsedURL, err := url.Parse(baseURL + service.ControlURL)

	if err != nil {
		return
	}

	client := soap.NewSOAPClient(*parsedURL)

	err = client.PerformAction(service.ServiceType, actionName, in, out)

	return err
}
