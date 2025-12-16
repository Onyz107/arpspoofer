//go:build linux

package sysctl

var sysctlValues = map[string]int{
	// Global master switch for IPv4 routing.
	// If this is 0, the kernel will NOT forward IPv4 packets at all,
	// regardless of any per-interface forwarding settings.
	"net.ipv4.ip_forward": 1,

	// Enable IPv4 forwarding on all currently existing interfaces.
	// This controls per-interface forwarding behavior but does NOT
	// override net.ipv4.ip_forward.
	"net.ipv4.conf.all.forwarding": 1,

	// Enable IPv6 forwarding on all currently existing interfaces.
	// Note: enabling this implicitly disables IPv6 Router Advertisement
	// acceptance (Linux assumes this host is a router).
	"net.ipv6.conf.all.forwarding": 1,

	// Enable loose reverse path filtering on all interfaces.
	//
	// Mode meanings:
	//   0 = Disabled
	//   1 = Strict (source IP must be reachable via the same interface
	//               the packet arrived on — breaks asymmetric routing
	//               and MITM / ARP spoofing in multi-NIC setups)
	//   2 = Loose  (source IP must be routable via ANY interface)
	//
	// Loose mode only drops packets with completely unroutable or
	// bogus source IPs and does NOT enforce interface symmetry.
	// This is safe for multi-NIC MITM setups.
	"net.ipv4.conf.all.rp_filter": 2,

	// Only reply to ARP requests if the target IP address is assigned
	// to the interface that received the ARP request.
	//
	// Without this, Linux may reply to ARP requests on the "wrong"
	// interface (ARP flux), causing hosts to learn incorrect
	// IP-to-MAC mappings in multi-NIC environments.
	"net.ipv4.conf.all.arp_ignore": 1,

	// Control how the kernel chooses the source IP address for
	// outgoing ARP requests.
	//
	// Mode 2 enforces that the ARP source IP must belong to the
	// outgoing interface, preventing Linux from advertising
	// IP addresses from other interfaces.
	//
	// This is critical in multi-NIC setups to prevent victims from
	// learning inconsistent or incorrect IP-to-MAC mappings.
	"net.ipv4.conf.all.arp_announce": 2,
}
