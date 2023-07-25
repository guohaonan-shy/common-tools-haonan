package network

import (
	"github.com/bytedance/sonic"
	"net"
	"os"
	"testing"
)

func TestLocalIPManager_Allocate(t *testing.T) {
	manager := &LocalIPManager{
		IpamDefaultStoragePath: "/home/guohaonan/code/src/common-tools-haonan/docker/network/ipam.json",
		IpamStorage:            map[string][]byte{},
	}

	var (
		ipamFile *os.File
	)

	_, err := os.Stat(manager.IpamDefaultStoragePath)
	if err != nil {
		if os.IsNotExist(err) {
			ipamFile, err = os.OpenFile(manager.IpamDefaultStoragePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				t.Fatal(err)
			}

			bytes, err := sonic.Marshal(manager)
			if err != nil {
				t.Fatal(err)
			}
			_, err = ipamFile.Write(bytes)
			if err != nil {
				t.Fatal(err)
			}
			ipamFile.Close()
		} else {
			t.Fatalf("dir stat failed, err:%s", err)
		}
	}

	_, ipRange, err := net.ParseCIDR("100.0.0.0/12")
	if err != nil {
		t.Error(err)
	}

	ip, err := manager.Allocate(ipRange)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ip)
}

func TestLocalIPManager_Release(t *testing.T) {
	manager := &LocalIPManager{
		IpamDefaultStoragePath: "/home/guohaonan/code/src/common-tools-haonan/docker/network/ipam.json",
		IpamStorage:            map[string][]byte{},
	}
	_, ipRange, err := net.ParseCIDR("100.0.0.0/12")
	if err != nil {
		t.Fatal(err)
	}

	ip := net.IPv4(100, 0, 0, 1)
	err = manager.Release(ipRange, &ip)
	if err != nil {
		t.Fatal(err)
	}
}
