package network

import (
	"fmt"
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

func (bridge *BridgeDriver) CreateNetwork(subnet *net.IPNet, networkName string) (*Network, error) {
	n := &Network{
		NetworkName: networkName,
		IPRange:     subnet,
		Driver:      bridge.Name(),
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
	cmds := fmt.Sprintf("-t nat -A POSTROUTING -s %s ! -o %s -j MASQUERADE", subnet.String(), networkName)
	output, err := exec.Command("iptables", strings.Split(cmds, " ")...).CombinedOutput()
	if err != nil {
		logrus.Errorf("[Bridge Driver] Create Network, exec iptables failed, err:%s, output:%s", err, string(output))
		return nil, err
	}

	return n, nil
}

// DeleteNetwork 删除网络
func (bridge *BridgeDriver) DeleteNetwork(network *Network) error {
	name := network.NetworkName
	br, err := netlink.LinkByName(name)
	if err != nil {
		return err
	}
	return netlink.LinkDel(br)
}

// Connect 容器连接, 创建容器网路通信组件veth 并设置为up
func (bridge *BridgeDriver) Connect(network *Network, endPoint *EndPoint) error {

	// 获取具体驱动网络信息 通过网络的唯一标识
	networkName := network.NetworkName
	br, err := netlink.LinkByName(networkName)
	if err != nil {
		return err
	}

	// 容器veth的信息
	la := netlink.NewLinkAttrs()
	la.Name = endPoint.ID[:5]
	la.MasterIndex = br.Attrs().Index

	// 创捷veth
	endPoint.Device = &netlink.Veth{
		LinkAttrs: la,
		PeerName:  "cif-" + endPoint.ID[:5],
	}

	// 创建接口
	if err = netlink.LinkAdd(endPoint.Device); err != nil {
		logrus.Errorf("[bridge] connect network failed, err:%s", err)
		return err
	}

	// 打开
	if err = netlink.LinkSetUp(endPoint.Device); err != nil {
		logrus.Errorf("[bridge] veth set up failed, err:%s", err)
		return err
	}

	return nil
}

// Disconnect 容器断连
func (bridge *BridgeDriver) Disconnect(network *Network, endPoint *EndPoint) error {
	return nil
}
