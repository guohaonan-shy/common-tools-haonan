package network

import (
	"github.com/bytedance/sonic"
	"net"
	"os"
)

var (
	ipAddressManager *LocalIPManager
)

func Init() {
	ipAddressManager = &LocalIPManager{}
}

type LocalIPManager struct {
	ipamDefaultStoragePath string
	ipamStorage            map[string][]byte
}

func (manager *LocalIPManager) Allocate(subnet *net.IPNet) (net.IP, error) {
	// 从宿主机内部读取ip分配信息
	if _, err := os.Stat(manager.ipamDefaultStoragePath); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	ipamFile, err := os.Open(manager.ipamDefaultStoragePath)
	if err != nil {
		return nil, err
	}
	var contentBytes []byte
	if _, err = ipamFile.Read(contentBytes); err != nil {
		return nil, err
	}

	if err = sonic.Unmarshal(contentBytes, &manager.ipamStorage); err != nil {
		return nil, err
	}

	ipPool := manager.ipamStorage

	// ip 分配
	subnetNum, ipNum := subnet.Mask.Size()
	if _, exist := ipPool[subnet.String()]; !exist {

	}
	return nil, nil
}

func (manager *LocalIPManager) Release(subnet *net.IPNet, ip *net.IP) error {
	return nil
}
