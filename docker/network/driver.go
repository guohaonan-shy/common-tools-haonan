package network

import (
	"net"
)

var (
	NetworkDrivers = make(map[string]Driver)
)

type Driver interface {
	// Name 驱动名称
	Name() string
	// CreateNetwork 根据指定驱动创建网络配置
	CreateNetwork(subnet *net.IPNet, name string) (*Network, error)
	// DeleteNetwork 删除网络
	DeleteNetwork(network *Network) error
	// Connect 容器连接
	Connect(network *Network, endPoint *EndPoint) error
	// Disconnect 容器断连
	Disconnect(network *Network, endPoint *EndPoint) error
}
