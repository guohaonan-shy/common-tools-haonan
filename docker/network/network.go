package network

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"path"
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

	network, err := driverInterface.CreateNetwork(subnet, networkName)
	if err != nil {
		logrus.Errorf("[CreateNetwork] driver:%v create failed, err:%s", reflect.ValueOf(driverInterface).Interface(), err)
		return err
	}

	return network.Dump(defaultNetworkPath)
}

func DeleteNetwork(networkName string) error {
	return nil
}
