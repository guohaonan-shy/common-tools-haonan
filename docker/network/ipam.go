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
	IpamDefaultStoragePath string            `json:"ipam_default_storage_path"`
	IpamStorage            map[string][]byte `json:"ipam_storage"`
}

func (manager *LocalIPManager) Allocate(subnet *net.IPNet) (net.IP, error) {

	var (
		contentBytes = make([]byte, 0)
		err          error
		ip           net.IP
	)
	// 从宿主机内部读取ip分配信息
	if _, err = os.Stat(manager.IpamDefaultStoragePath); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	if contentBytes, err = os.ReadFile(manager.IpamDefaultStoragePath); err != nil {
		return nil, err
	}

	if err = sonic.Unmarshal(contentBytes, &manager); err != nil {
		return nil, err
	}

	ipPool := manager.IpamStorage

	// ip 分配
	// 初始化bitmap
	subnetNum, total := subnet.Mask.Size()
	if _, exist := ipPool[subnet.String()]; !exist {
		avaiableIpNum := 1 << (total - subnetNum)
		byteNum, rest := avaiableIpNum/8, avaiableIpNum%8

		pool := bytes.Repeat([]byte{0x00}, byteNum)
		// rest非0，需要添加rest个可用位置
		if rest != 0 {
			pool = append(pool, 0xff>>rest)
		}

		ipPool[subnet.String()] = pool
	}

	// 分配
	subnetIpPool := ipPool[subnet.String()]

	for i := range subnetIpPool {
		bit, base := subnetIpPool[i], byte(0x80)
		isUnusedExisted, bitNum := false, 0
		for j := 0; j < 8; j++ {
			re := bit & (base >> j)
			if re == base>>j {
				// 该位已经被使用
				continue
			}
			isUnusedExisted = true
			bitNum = j
			// 修改为已占用
			subnetIpPool[i] = bit | (base >> j)
			break
		}

		if !isUnusedExisted {
			continue
		}

		// 具体位数
		index := i*8 + bitNum + 1
		ip = subnet.IP

		for t := uint(4); t > 0; t -= 1 {
			[]byte(ip)[4-t] += uint8(index >> ((t - 1) * 8))
		}
		//ip[3] += 1
		break
	}

	ipPool[subnet.String()] = subnetIpPool

	manager.IpamStorage = ipPool

	// 重新dump
	if _, err = os.Stat(manager.IpamDefaultStoragePath); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	newfile, err := os.OpenFile(manager.IpamDefaultStoragePath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
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

	return ip, nil
}

func (manager *LocalIPManager) Release(subnet *net.IPNet, ip *net.IP) error {
	// 从宿主机内部读取ip分配信息
	if _, err := os.Stat(manager.IpamDefaultStoragePath); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	ipamFile, err := os.Open(manager.IpamDefaultStoragePath)
	if err != nil {
		return err
	}
	var contentBytes []byte
	if _, err = ipamFile.Read(contentBytes); err != nil {
		return err
	}

	if err = sonic.Unmarshal(contentBytes, &manager); err != nil {
		return err
	}

	ipPool := manager.IpamStorage

	c := 0
	ipFor4 := ip.To4()
	ipFor4[3] -= 1
	for t := uint(4); t > 0; t -= 1 {
		c += int(ipFor4[t-1] - subnet.IP[t-1]<<((4-t)*8))
	}

	ipPool[subnet.String()][c] = 0

	manager.IpamStorage = ipPool

	// 重新dump
	if _, err = os.Stat(manager.IpamDefaultStoragePath); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	newfile, err := os.OpenFile(manager.IpamDefaultStoragePath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
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
