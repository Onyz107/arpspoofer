package handle

import (
	"net"

	"github.com/Onyz107/arpspoofer/internal/arp"
	"github.com/Onyz107/arpspoofer/internal/logger"
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
		logger.Logger.Info(pkt.String())
	}
	return h.handle.WritePacketData(pkt.Bytes())
}

func (h *Handle) NewPacketSource() *gopacket.PacketSource {
	return gopacket.NewPacketSource(h.handle, h.handle.LinkType())
}

func (h *Handle) Close() {
	h.handle.Close()
}
