//go:build darwin

package sysctl

var sysctlValues = map[string]int{
	// Global master switch for IPv4 routing.
	// If this is 0, the kernel will NOT forward IPv4 packets at all,
	// regardless of any per-interface forwarding settings.
	"net.inet.ip.forwarding": 1,

	// Enable IPv6 forwarding on all currently existing interfaces.
	"net.inet6.ip6.forwarding": 1,
}
