package driver

import (
	"github.com/common-tools-haonan/docker/network"
	"net"
)

var (
	NetworkDrivers = map[string]Driver{}
)

type Driver interface {
	// Name 驱动名称
	Name() string
	// CreateNetwork 根据指定驱动创建网络配置
	CreateNetwork(subnet *net.IPNet, name string) (*network.Network, error)
	// DeleteNetwork 删除网络
	DeleteNetwork(network *network.Network) error
	// Connect 容器连接
	Connect(network *network.Network, endPoint *network.EndPoint) error
	// Disconnect 容器断连
	Disconnect(network *network.Network, endPoint *network.EndPoint) error
}
