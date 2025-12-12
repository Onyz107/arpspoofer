package hwid

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/Onyz107/arpspoofer/handle"
	"github.com/Onyz107/arpspoofer/internal/arp"
	"github.com/google/gopacket/layers"
)

func GetFromIP(ctx context.Context, ifaceHandle *handle.Handle, ip net.IP, verbose bool) (net.HardwareAddr, error) {
	inCtx, cancel := context.WithTimeout(ctx, timeoutTime)
	defer cancel()

	pkt, err := arp.BuildPkt(ifaceHandle.HWID, ifaceHandle.IP, ifaceHandle.HWID, ip, broadcastMAC, arp.ARPRequest)
	if err != nil {
		return nil, errors.Join(ErrBuildPkt, err)
	}

	src := ifaceHandle.NewPacketSource()
	packets := src.Packets()

	tick := time.NewTicker(tickerTime)
	defer tick.Stop()

	for {
		if err := ifaceHandle.WritePacketData(pkt, verbose); err != nil {
			return nil, errors.Join(ErrWritePacketData, err)
		}

		select {
		case <-inCtx.Done():
			return nil, ErrTimeout

		case packet, ok := <-packets:
			if !ok {
				return nil, ErrFindHWID
			}

			arpLayer := packet.Layer(layers.LayerTypeARP)
			if arpLayer == nil {
				continue
			}
			arpPkt := arpLayer.(*layers.ARP)

			if arpPkt.Operation == layers.ARPReply && net.IP(arpPkt.SourceProtAddress).Equal(ip) {
				return net.HardwareAddr(arpPkt.SourceHwAddress), nil
			}

		case <-tick.C:
			continue
		}
	}
}
