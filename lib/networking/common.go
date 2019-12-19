package networking

type BridgeMasterInterfaceConfig struct {
	Device  string `json:"device"`
	Enabled bool   `json:"enabled"`
}

type IP4Config struct {
	Enabled bool   `json:"enabled"`
	DHCP    bool   `json:"dhcp"`
	Address string `json:"address,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Vlan    uint16 `json:"vlan"`
}

type NetworkConfig struct {
	ID              string
	Name            string
	Type            BridgeDriver
	Enabled         bool                         `json:"enabled"`
	MasterInterface *BridgeMasterInterfaceConfig `json:"masterInterface"`
	IPV4            *IP4Config                   `json:"ipv4"`
}

type BridgeDriver string

const (
	LinuxBridgeDriver BridgeDriver = "linux"
	OvsBridgeDriver   BridgeDriver = "ovs"
)

type NetworkBridge interface {
	GetId() string
	GetName() string
	SetName(name string)
	CreateInterface(vmid string, index uint, vlan uint16) (NetworkInterface, error)
	GetInterface(interfaceId string) (NetworkInterface, error)
	DestroyInterface(interfaceId string)
}

type NetworkInterface interface {
	GetId() string
	Enable() error
	Disable() error
}
