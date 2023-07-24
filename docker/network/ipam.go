package network

import (
	"bytes"
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
	// 初始化bitmap
	subnetNum, total := subnet.Mask.Size()
	if _, exist := ipPool[subnet.String()]; !exist {
		avaiableIpNum := 1 << (total - subnetNum)
		byteNum, rest := avaiableIpNum/8, avaiableIpNum%8

		pool := bytes.Repeat([]byte{0x00}, byteNum)
		pool = append(pool, 0xff>>rest)

		ipPool[subnet.String()] = pool

	}

	// 分配
	subnetIpPool := ipPool[subnet.String()]

	for i := range subnetIpPool {
		if subnetIpPool[i] == 0 {

			subnetIpPool[i] = 1
			ip := subnet.IP

			for t := uint(4); t > 0; t -= 1 {
				[]byte(ip)[4-t] = uint8(i >> ((t - 1) * 8))
			}
			ip[3] += 1
			break
		}
	}

	ipPool[subnet.String()] = subnetIpPool

	manager.ipamStorage = ipPool

	// 重新dump
	if _, err = os.Stat(manager.ipamDefaultStoragePath); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	newfile, err := os.OpenFile(manager.ipamDefaultStoragePath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	var content []byte
	if content, err = sonic.Marshal(manager); err != nil {
		return nil, err
	}

	if _, err = newfile.Write(content); err != nil {
		return nil, err
	}

	return nil, nil
}

func (manager *LocalIPManager) Release(subnet *net.IPNet, ip *net.IP) error {
	// 从宿主机内部读取ip分配信息
	if _, err := os.Stat(manager.ipamDefaultStoragePath); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	ipamFile, err := os.Open(manager.ipamDefaultStoragePath)
	if err != nil {
		return err
	}
	var contentBytes []byte
	if _, err = ipamFile.Read(contentBytes); err != nil {
		return err
	}

	if err = sonic.Unmarshal(contentBytes, &manager.ipamStorage); err != nil {
		return err
	}

	ipPool := manager.ipamStorage

	c := 0
	ipFor4 := ip.To4()
	ipFor4[3] -= 1
	for t := uint(4); t > 0; t -= 1 {
		c += int(ipFor4[t-1] - subnet.IP[t-1]<<((4-t)*8))
	}

	ipPool[subnet.String()][c] = 0

	manager.ipamStorage = ipPool

	// 重新dump
	if _, err = os.Stat(manager.ipamDefaultStoragePath); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	newfile, err := os.OpenFile(manager.ipamDefaultStoragePath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	var content []byte
	if content, err = sonic.Marshal(manager); err != nil {
		return err
	}

	if _, err = newfile.Write(content); err != nil {
		return err
	}

	return nil
}
