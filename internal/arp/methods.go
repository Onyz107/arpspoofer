package arp

import (
	"fmt"
	"net"
	"sync"

	"github.com/google/gopacket"
)

type ARPPkt struct {
	bytes          []byte
	srcMAC, dstMAC net.HardwareAddr
	srcIP, dstIP   net.IP
	operation      ARPType
}

func (a *ARPPkt) Bytes() []byte {
	return a.bytes
}

func (a *ARPPkt) String() string {
	if a.operation == ARPRequest {
		return fmt.Sprintf("ARP Request: Who has %s? Tell %s at %s --> %s", a.dstIP.To4().String(), a.srcIP.To4().String(),
			a.srcMAC.String(), a.dstMAC.String())
	} else {
		return fmt.Sprintf("ARP Reply: %s is at %s --> %s at %s", a.srcIP.To4().String(), a.srcMAC.String(),
			a.dstIP.To4().String(), a.dstMAC.String())
	}
}

var bufPool = sync.Pool{
	New: func() any {
		return gopacket.NewSerializeBuffer()
	},
}
