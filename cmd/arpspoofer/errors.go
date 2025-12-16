package main

import "errors"

var (
	ErrInvalidTargetIP       = errors.New("invalid target IP address")
	ErrInvalidHostIP         = errors.New("invalid host IP address")
	ErrInvalidInterface      = errors.New("invalid network interface")
	ErrOpenInterface         = errors.New("failed to open network interface")
	ErrStartSpoofing         = errors.New("failed to start ARP spoofing")
	ErrInvalidSysctlSettings = errors.New("invalid sysctl settings for ARP spoofing")
	ErrInterfaceNotFound     = errors.New("network interface not found")
	ErrInterfaceDown         = errors.New("network interface is down")
	ErrDialFailed            = errors.New("failed to create outbound UDP connection")
	ErrInterfaceListFailed   = errors.New("failed to list network interfaces")
	ErrIPNotOnInterface      = errors.New("interface does not own the outbound IP")
)
