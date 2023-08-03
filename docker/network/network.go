package network

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/common-tools-haonan/docker/container"
	"github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netns"
	"net"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
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
		logrus.Errorf("[network Load]failed:%s", err)
		return err
	}

	return sonic.Unmarshal(bytes, &network)
}

func (network *Network) Remove() error {
	networkPath := path.Join(defaultNetworkPath, "/", network.NetworkName)
	_, err := os.Stat(networkPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			logrus.Errorf("[Network Remove] stat(which is used to check whether the file is exised) failed, err:%s", err)
			return err
		}
	}

	return os.RemoveAll(networkPath)
}

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
	network, ok := networkMapping[networkName]
	if !ok {
		return errors.New(fmt.Sprintf("network:%s not existed", networkName))
	}

	ipRange := network.IPRange
	err := ipAddressManager.Release(ipRange, &ipRange.IP)
	if err != nil {
		return err
	}

	var (
		driver        Driver
		isDriverExist bool
	)

	driver, isDriverExist = NetworkDrivers[network.Driver]
	if !isDriverExist {
		return errors.New(fmt.Sprintf("driver:%s not init", network.Driver))
	}

	if err = driver.DeleteNetwork(network); err != nil {
		return err
	}

	return network.Remove()
}

func ListAllNetwork() ([]*Network, error) {
	// 读取默认目录
	if _, err := os.Stat(defaultNetworkPath); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(defaultNetworkPath, 0644)
		} else {
			logrus.Error(err)
			return nil, err
		}
	}

	networks := make([]*Network, 0)
	err := filepath.Walk(defaultNetworkPath, func(networkPath string, file os.FileInfo, err error) error {
		// 底层默认执行一次walkfc, 避免报错，这步过滤
		if file.IsDir() {
			return nil
		}

		if file.Name() == "ipam_config.json" {
			return nil
		}

		_, networkFileName := path.Split(networkPath)
		network := &Network{
			NetworkName: networkFileName,
		}

		if err = network.Load(networkPath); err != nil {
			return err
		}

		networks = append(networks, network)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return networks, nil
}

func Connect(networkName string, portMapping string, containerId string) error {
	network, ok := networkMapping[networkName]
	if !ok {
		return errors.New(fmt.Sprintf("network:%s not existed", networkName))
	}

	// 接入容器的网络ip分配
	ipRange := network.IPRange
	ip, err := ipAddressManager.Allocate(ipRange)
	if err != nil {
		return err
	}

	endpoint := &EndPoint{
		ID:          fmt.Sprintf("%s-%s", containerId, networkName),
		IPAddress:   &ip,
		Network:     network,
		PortMapping: strings.Split(portMapping, ":"),
	}

	var (
		driver        Driver
		isDriverExist bool
	)

	driver, isDriverExist = NetworkDrivers[network.Driver]
	if !isDriverExist {
		return errors.New(fmt.Sprintf("driver:%s not init", network.Driver))
	}

	// 容器网络设备veth创建
	if err = driver.Connect(network, endpoint); err != nil {
		return err
	}

	// 配置ip, route
	if err = configInterfaceIpAndRoute(endpoint, containerId); err != nil {
		return err
	}

	// port
	if err = configPortMapping(endpoint); err != nil {
		return err
	}
	return nil
}

func configInterfaceIpAndRoute(endpoint *EndPoint, containerId string) (err error) {

	network := endpoint.Network

	peerLink, err := netlink.LinkByName(endpoint.Device.PeerName)
	if err != nil {
		return fmt.Errorf("fail config endpoint: %v", err)
	}

	containerInfo, err := container.GetSpecificContainers(containerId)

	defer enterContainerNetns(&peerLink, containerInfo)()

	if err = setInterfaceIp(endpoint.Device.PeerName, network.IPRange.String()); err != nil {
		return err
	}

	if err = setInterfaceUp(endpoint.Device.PeerName); err != nil {
		return err
	}

	if err = setInterfaceUp("lo"); err != nil {
		return err
	}

	interfaceDev, err := netlink.LinkByName(endpoint.Device.PeerName)
	if err != nil {
		logrus.Errorf("[configInterfaceIpAndRoute] interface:%s find failed, err:%s", endpoint.Device.PeerName, err)
		return err
	}
	_, cidr, _ := net.ParseCIDR("0.0.0.0/0")
	defaultRoute := &netlink.Route{
		LinkIndex: interfaceDev.Attrs().Index,
		Gw:        network.IPRange.IP,
		Dst:       cidr,
	}

	if err = netlink.RouteAdd(defaultRoute); err != nil {
		logrus.Errorf("[configInterfaceIpAndRoute] route add failed, err:%s", err)
		return err
	}

	return nil
}

func setInterfaceIp(name string, ipNet string) error {
	interfaceDev, err := netlink.LinkByName(name)
	if err != nil {
		logrus.Errorf("[setInterfaceIp] interface:%s find failed, err:%s", name, err)
		return err
	}

	_, subnet, err := net.ParseCIDR(ipNet)
	if err != nil {
		logrus.Errorf("[serInterfaceIp] ipnet:%s parse failed, err:%s", ipNet, err)
		return err
	}

	addr := &netlink.Addr{
		IPNet: subnet,
		Label: "",
		Flags: 0,
		Scope: 0,
	}

	if err = netlink.AddrAdd(interfaceDev, addr); err != nil {
		logrus.Errorf("[setInterfance] addr add failed, err:%s", err)
		return err
	}
	return nil
}

func setInterfaceUp(name string) error {
	interfaceDev, err := netlink.LinkByName(name)
	if err != nil {
		logrus.Errorf("[setInterfaceUp] interface:%s find failed, err:%s", name, err)
		return err
	}

	if err = netlink.LinkSetUp(interfaceDev); err != nil {
		logrus.Errorf("[setInterfaceUp] Link set up failed, err:%s", err)
		return err
	}
	return nil
}

func enterContainerNetns(enLink *netlink.Link, cinfo *container.ContainerInfo) func() {
	f, err := os.OpenFile(fmt.Sprintf("/proc/%s/ns/net", cinfo.Pid), os.O_RDONLY, 0)
	if err != nil {
		logrus.Errorf("error get container net namespace, %v", err)
	}

	nsFD := f.Fd()
	runtime.LockOSThread()

	// 修改veth peer 另外一端移到容器的namespace中
	if err = netlink.LinkSetNsFd(*enLink, int(nsFD)); err != nil {
		logrus.Errorf("error set link netns , %v", err)
	}

	// 获取当前的网络namespace
	origns, err := netns.Get()
	if err != nil {
		logrus.Errorf("error get current netns, %v", err)
	}

	// 设置当前进程到新的网络namespace，并在函数执行完成之后再恢复到之前的namespace
	if err = netns.Set(netns.NsHandle(nsFD)); err != nil {
		logrus.Errorf("error set netns, %v", err)
	}
	return func() {
		netns.Set(origns)
		origns.Close()
		runtime.UnlockOSThread()
		f.Close()
	}
}

func configPortMapping(ep *EndPoint) error {
	for _, pm := range ep.PortMapping {
		portMapping := strings.Split(pm, ":")
		if len(portMapping) != 2 {
			logrus.Errorf("port mapping format error, %v", pm)
			continue
		}
		iptablesCmd := fmt.Sprintf("-t nat -A PREROUTING -p tcp -m tcp --dport %s -j DNAT --to-destination %s:%s",
			portMapping[0], ep.IPAddress.String(), portMapping[1])
		cmd := exec.Command("iptables", strings.Split(iptablesCmd, " ")...)
		//err := cmd.Run()
		output, err := cmd.Output()
		if err != nil {
			logrus.Errorf("iptables Output, %v", output)
			continue
		}
	}
	return nil
}
