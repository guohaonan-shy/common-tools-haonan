package driver

import (
	"fmt"
	"github.com/common-tools-haonan/docker/network"
	"github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
	"net"
	"os/exec"
	"strings"
)

type BridgeDriver struct {
}

func (bridge *BridgeDriver) Name() string {
	return "bridge"
}

func (bridge *BridgeDriver) CreateNetwork(subnet *net.IPNet, networkName string) (*network.Network, error) {
	n := &network.Network{
		NetworkName: networkName,
		IPRange:     subnet,
	}

	// 1. 创建网桥实例
	// 1.1 先通过name判断网络接口是否存在，若没有
	_, err := net.InterfaceByName(networkName)
	if err != nil && !strings.Contains(err.Error(), "no such network interface") {
		return nil, err
	}

	// 1.2 设置link属性，并创建网桥
	la := netlink.NewLinkAttrs()
	la.Name = networkName

	bridgeIns := &netlink.Bridge{
		LinkAttrs: la,
	}

	err = netlink.LinkAdd(bridgeIns)
	if err != nil {
		logrus.Errorf("[Bridge Driver] Create Network, Create Bridge failed, err:%s", err)
		return nil, err
	}

	// 2. 设置网桥网段， 即宿主机将子网网段的请求路由到创建的该网桥上
	link, err := netlink.LinkByName(networkName)
	if err != nil {
		return nil, err
	}

	err = netlink.AddrAdd(link, &netlink.Addr{
		IPNet: subnet,
		Label: "",
		Flags: 0,
		Scope: 0,
	})

	if err != nil {
		logrus.Errorf("[Bridge Driver] Create Network, Add Route failed, err:%s", err)
		return nil, err
	}

	// 3. 打开网桥， ip link set bro0 up
	link, err = netlink.LinkByName(networkName)
	if err != nil {
		return nil, err
	}

	err = netlink.LinkSetUp(link)
	if err != nil {
		logrus.Errorf("[Bridge Driver] Create Network, Set Bridge From Down to Up failed, err:%s", err)
		return nil, err
	}
	// 4. SNAT规则
	// iptables -t nat -A POSTROUTING -s subnet ! -o name -j "MASQUERADE"
	output, err := exec.Command("iptables",
		fmt.Sprintf("-t nat -A POSTROUTING -s %s ! -o %s -j MASQUERATE", subnet.String(), networkName),
	).CombinedOutput()
	if err != nil {
		logrus.Errorf("[Bridge Driver] Create Network, exec iptables failed, err:%s, output:%s", string(output))
		return nil, err
	}

	return n, nil
}

// DeleteNetwork 删除网络
func (bridge *BridgeDriver) DeleteNetwork(network *network.Network) error {
	name := network.NetworkName
	br, err := netlink.LinkByName(name)
	if err != nil {
		return err
	}
	return netlink.LinkDel(br)
}

// Connect 容器连接
func (bridge *BridgeDriver) Connect(network *network.Network, endPoint *network.EndPoint) error {
	return nil
}

// Disconnect 容器断连
func (bridge *BridgeDriver) Disconnect(network *network.Network, endPoint *network.EndPoint) error {
	return nil
}
