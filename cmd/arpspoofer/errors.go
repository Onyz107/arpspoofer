package main

import "errors"

var ErrInvalidTargetIP = errors.New("invalid target IP address")
var ErrInvalidHostIP = errors.New("invalid host IP address")
var ErrInvalidInterface = errors.New("invalid network interface")
var ErrOpenInterface = errors.New("failed to open network interface")
var ErrStartSpoofing = errors.New("failed to start ARP spoofing")
var ErrInvalidSysctlSettings = errors.New("invalid sysctl settings for ARP spoofing")
