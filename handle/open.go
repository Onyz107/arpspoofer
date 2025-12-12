package handle

import (
	"errors"
	"net"

	"github.com/google/gopacket/pcap"
)

func Open(iface string) (*Handle, error) {
	handle, err := pcap.OpenLive(iface, 65536, true, pcap.BlockForever)
	if err != nil {
		return nil, errors.Join(ErrOpenInterface, err)
	}

	hwaddr, err := getHWID(iface)
	if err != nil {
		handle.Close()
		return nil, errors.Join(ErrGetHWID, err)
	}

	ipaddr, err := getIPAddr(iface)
	if err != nil {
		handle.Close()
		return nil, errors.Join(ErrGetIPAddr, err)
	}

	return &Handle{handle: handle, IP: ipaddr, HWID: hwaddr}, nil
}

func getHWID(iface string) (net.HardwareAddr, error) {
	netIface, err := net.InterfaceByName(iface)
	if err != nil {
		return nil, errors.Join(ErrHWIDNotFound, err)
	}

	return netIface.HardwareAddr, nil
}

func getIPAddr(iface string) (net.IP, error) {
	netIface, err := net.InterfaceByName(iface)
	if err != nil {
		return nil, errors.Join(ErrIfaceName, err)
	}

	addrs, err := netIface.Addrs()
	if err != nil {
		return nil, errors.Join(ErrIfaceAddrs, err)
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP, nil
			}
		}
	}

	return nil, ErrNoIPAddr
}
