package arp

import (
	"errors"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func BuildPkt(ifaceHWID net.HardwareAddr, srcIP net.IP, srcMAC net.HardwareAddr, dstIP net.IP, dstMAC net.HardwareAddr, operation ARPType) (*ARPPkt, error) {
	eth := &layers.Ethernet{
		SrcMAC:       ifaceHWID,
		DstMAC:       dstMAC,
		EthernetType: layers.EthernetTypeARP,
	}

	arp := &layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         uint16(operation),
		SourceHwAddress:   srcMAC,
		SourceProtAddress: []byte(srcIP.To4()),
		DstHwAddress:      dstMAC,
		DstProtAddress:    []byte(dstIP.To4()),
	}

	buf := bufPool.Get().(gopacket.SerializeBuffer)
	defer func() {
		buf.Clear()
		bufPool.Put(buf)
	}()

	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	if err := gopacket.SerializeLayers(buf, opts, eth, arp); err != nil {
		return nil, errors.Join(ErrSerializeLayers, err)
	}

	out := append([]byte(nil), buf.Bytes()...)

	pkt := &ARPPkt{
		bytes:     out,
		srcMAC:    srcMAC,
		dstMAC:    dstMAC,
		srcIP:     srcIP,
		dstIP:     dstIP,
		operation: operation,
	}

	return pkt, nil
}
