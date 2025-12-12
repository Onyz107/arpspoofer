package handle

import (
	"log"
	"net"

	"github.com/Onyz107/arpspoofer/internal/arp"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type Handle struct {
	handle *pcap.Handle
	IP     net.IP
	HWID   net.HardwareAddr
}

func (h *Handle) WritePacketData(pkt *arp.ARPPkt, verbose bool) error {
	if verbose {
		log.Println(pkt.GetInfo())
	}
	return h.handle.WritePacketData(pkt.Bytes())
}

func (h *Handle) NewPacketSource() *gopacket.PacketSource {
	return gopacket.NewPacketSource(h.handle, h.handle.LinkType())
}

func (h *Handle) Close() {
	h.handle.Close()
}
