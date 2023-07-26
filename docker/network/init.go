package network

import "github.com/sirupsen/logrus"

var (
	ipAddressManager *LocalIPManager
	networkMapping   = make(map[string]*Network)
	driverMapping    = make(map[string]Driver)
)

func Init() {
	ipAddressManager = &LocalIPManager{
		IpamDefaultStoragePath: defaultNetworkPath + "/ipam_config.json",
		IpamStorage:            make(map[string][]byte),
	}

	networks, err := ListAllNetwork()
	if err != nil {
		logrus.Fatal(err)
	}

	for i, network := range networks {
		if _, exist := networkMapping[network.NetworkName]; !exist {
			networkMapping[network.NetworkName] = networks[i]
		}
	}

	// bridge
	bridge := &BridgeDriver{}
	driverMapping[bridge.Name()] = bridge

}
