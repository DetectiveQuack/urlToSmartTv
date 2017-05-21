package device

// TV xml struct1
type TV struct {
	SpecVersion Version `xml:"specVersion"`
	DeviceInfo  Device  `xml:"device"`
}

// Version from samsung xml
type Version struct {
	Major int `xml:"major"`
	Minor int `xml:"minor"`
}

// Device struct holds device information
type Device struct {
	Name     string    `xml:"friendlyName"`
	Services []Service `xml:"serviceList>service"`
}

// Service struct holds all available services
type Service struct {
	ServiceID   string `xml:"serviceId"`
	ServiceType string `xml:"serviceType"`
	ControlURL  string `xml:"controlURL"`
}

// OutAction sdsd
type OutAction struct {
	ConnectionIDs string
}

// InAction sdsd
type InAction struct{}

// PrepareForConnectionInArgs holds all input args for ConnectionManager#PrepareForConnection
type PrepareForConnectionInArgs struct {
	RemoteProtocolInfo    string `soap:"RemoteProtocolInfo"`
	PeerConnectionManager string `soap:"PeerConnectionManager"`
	PeerConnectionID      string `soap:"PeerConnectionID"`
	Direction             string `soap:"Direction"`
}

// PrepareForConnectionOutArgs holds all output args for ConnectionManager#PrepareForConnection
type PrepareForConnectionOutArgs struct {
	ConnectionID  string `soap:"ConnectionID"`
	AVTransportID string `soap:"AVTransportID"`
	RcsID         string `soap:"RcsID"`
}

// SetAVTransportURIInArgs struct holds input args got AVTransport#SetAVTransportURI
type SetAVTransportURIInArgs struct {
	InstanceID         string `soap:"InstanceID"`
	CurrentURI         string `soap:"CurrentURI"`
	CurrentURIMetaData string `soap:"CurrentURIMetaData"`
}

// SetAVTransportURIOutArgs struct holds output args got AVTransport#SetAVTransportURI
type SetAVTransportURIOutArgs struct{}

// PlayInArgs struct holds play args for AVTransport#Play
type PlayInArgs struct {
	InstanceID string
	Speed      string
}
