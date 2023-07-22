package network

import (
	"github.com/vishvananda/netlink"
	"net"
)

type EndPoint struct {
	ID          string
	Device      *netlink.Veth
	IPAddress   *net.IP
	MacAddress  *net.HardwareAddr
	PortMapping []string
	Network     *Network
}
