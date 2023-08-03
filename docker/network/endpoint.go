package network

import (
	"github.com/vishvananda/netlink"
	"net"
)

// EndPoint ghndocker server 针对容器网络通信的统一抽象
type EndPoint struct {
	ID          string
	Device      *netlink.Veth
	IPAddress   *net.IP
	MacAddress  *net.HardwareAddr
	PortMapping []string
	Network     *Network
}
