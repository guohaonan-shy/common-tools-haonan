package network

import "net"

var (
	ipAddressManager *LocalIPManager
)

func Init() {
	ipAddressManager = &LocalIPManager{}
}

type LocalIPManager struct {
}

func (manager *LocalIPManager) Allocate(subnet *net.IPNet) (net.IP, error) {
	return nil, nil
}

func (manager *LocalIPManager) Release(subnet *net.IPNet, ip *net.IP) error {
	return nil
}
