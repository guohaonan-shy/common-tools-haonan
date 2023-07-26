package network

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/common-tools-haonan/docker/container"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"path"
	"path/filepath"
	"reflect"
)

const (
	defaultNetworkPath = "/home/guohaonan/ghndocker/network"
)

type Network struct {
	NetworkName string     `json:"network_name"`
	Driver      string     `json:"driver"`
	IPRange     *net.IPNet `json:"ip_range"`
}

func (network *Network) Dump(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if mkDirErr := os.MkdirAll(dir, 0644); mkDirErr != nil {
				logrus.Errorf("[Network Dump] mk dir:%s failed, err:%s", dir, mkDirErr)
				return mkDirErr
			}
		} else {
			logrus.Errorf("[Network Dump]Dump failed, err:%s", err)
			return err
		}
	}

	networkPath := path.Join(defaultNetworkPath, "/", network.NetworkName)
	networkFile, err := os.OpenFile(networkPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		logrus.Errorf("[Network Dump] open files failed, err:%s", err)
		return err
	}

	defer networkFile.Close()

	bytes, err := sonic.Marshal(network)
	if err != nil {
		logrus.Errorf("[Network Dump] marshal failed, err:%s", err)
		return err
	}

	_, err = networkFile.Write(bytes)
	if err != nil {
		logrus.Errorf("[Network Dump] write failed, err:%s", err)
		return err
	}

	return nil
}

func (network *Network) Load(fileName string) error {

	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	return sonic.Unmarshal(bytes, &network)
}

//func (network *Network) Remove(networkName string) error {
//	networkPath := path.Join(defaultNetworkPath, "/", network.NetworkName)
//	_, err := os.Stat(networkPath)
//	if err != nil {
//		logrus.Errorf("[Network Remove] stat(which is used to check whether the file is exised) failed, err:%s", err)
//		return err
//	}
//
//	return os.RemoveAll(networkPath)
//}

func CreateNetwork(networkName, driverName string, subnet string) error {
	_, ipRange, err := net.ParseCIDR(subnet)
	if err != nil {
		logrus.Errorf("[CreateNetwork] create local network failed, err:%s", err)
		return err
	}

	// 网关地址
	gatewayIP, err := ipAddressManager.Allocate(ipRange)
	if err != nil {
		return err
	}
	ipRange.IP = gatewayIP

	var (
		driverInterface Driver
		isExist         bool
	)

	driverInterface, isExist = NetworkDrivers[driverName]
	if !isExist {
		logrus.Errorf("[CreateWork] driver:%s doesn't implement", driverName)
		return errors.New(fmt.Sprintf("driver:%s doesn't implement", driverName))
	}

	network, err := driverInterface.CreateNetwork(ipRange, networkName)
	if err != nil {
		logrus.Errorf("[CreateNetwork] driver:%v create failed, err:%s", reflect.ValueOf(driverInterface).Interface(), err)
		return err
	}

	return network.Dump(defaultNetworkPath)
}

func DeleteNetwork(networkName string) error {
	return nil
}

func ListAllNetwork() ([]*Network, error) {
	// 读取默认目录
	if _, err := os.Stat(defaultNetworkPath); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(defaultNetworkPath, 0644)
		} else {
			return nil, err
		}
	}

	networks := make([]*Network, 0)
	err := filepath.Walk(defaultNetworkPath, func(networkPath string, file os.FileInfo, err error) error {
		// 底层默认执行一次walkfc, 避免报错，这步过滤
		if file.IsDir() {
			return nil
		}

		_, networkFileName := path.Split(networkPath)
		network := &Network{
			NetworkName: networkFileName,
		}

		if err = network.Load(networkFileName); err != nil {
			return err
		}

		networks = append(networks, network)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func Connect(networkName string, container *container.ContainerInfo) error {
	network, ok := networkMapping[networkName]
	if !ok {
		return errors.New(fmt.Sprintf("network:%s not existed", networkName))
	}

	ipRange := network.IPRange
	ip, err := ipAddressManager.Allocate(ipRange)
	if err != nil {
		return err
	}

	endpoint := &EndPoint{
		ID:        fmt.Sprintf("%s-%s", container.Id, networkName),
		IPAddress: &ip,
		Network:   network,
	}

}
